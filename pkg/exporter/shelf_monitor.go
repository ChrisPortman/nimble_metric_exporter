package exporter

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/ChrisPortman/nimble_metric_exporter/internal/models"
	"github.com/ChrisPortman/nimble_metric_exporter/pkg/client"
)

type shelfLocationLookup struct {
	// Store a mapping of shelf IDs to human friendly location references
	shelfLocations     map[string]string
	shelfLocationsLock sync.RWMutex
}

type ShelfMetrics struct {
	client *client.NimbleClient
	logger *slog.Logger

	overallPSU  metric.Int64ObservableGauge
	overallTemp metric.Int64ObservableGauge
	overallFan  metric.Int64ObservableGauge

	sensorStatus metric.Int64ObservableGauge
	sensorValue  metric.Int64ObservableGauge

	shelfLocations *shelfLocationLookup
}

func NewShelfMetrics(client *client.NimbleClient, meter metric.Meter, logger *slog.Logger) (ShelfMetrics, error) {
	var err error

	log := slog.New(slog.DiscardHandler)
	if logger != nil {
		log = logger
	}

	metrics := ShelfMetrics{
		client:         client,
		logger:         log,
		shelfLocations: &shelfLocationLookup{},
	}

	metrics.overallPSU, err = meter.Int64ObservableGauge(
		"nimble.shelf.psus.ok",
		metric.WithDescription("Shelf powersupplies are OK (= 1)"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.shelf.psus.ok: %w", err)
	}

	metrics.overallFan, err = meter.Int64ObservableGauge(
		"nimble.shelf.fans.ok",
		metric.WithDescription("Shelf fans are OK (= 1)"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.shelf.fans.ok: %w", err)
	}

	metrics.overallTemp, err = meter.Int64ObservableGauge(
		"nimble.shelf.temp.ok",
		metric.WithDescription("Indicates temperature is OK overall (= 1)"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.shelf.temp.ok: %w", err)
	}

	metrics.sensorStatus, err = meter.Int64ObservableGauge(
		"nimble.shelf.sensor.ok",
		metric.WithDescription("Indicates sensor is OK (= 1)"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.shelf.sensor.ok: %w", err)
	}

	metrics.sensorValue, err = meter.Int64ObservableGauge(
		"nimble.shelf.sensor.value",
		metric.WithDescription("Value of sensor"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.shelf.sensor.value: %w", err)
	}

	if err := metrics.Register(meter); err != nil {
		return metrics, err
	}

	return metrics, nil
}

func (m *ShelfMetrics) Register(meter metric.Meter) error {
	ternary := func(in bool) int64 {
		if in {
			return 1
		}

		return 0
	}

	_, err := meter.RegisterCallback(
		func(ctx context.Context, observer metric.Observer) error {
			m.logger.Debug("loading shelf metrics")

			m.updateShelfLocations(ctx)

			shelvestates, err := m.getShelfStates(ctx)
			if err != nil {
				m.logger.Error("error retrieving shelf data", slog.String("error", err.Error()))

				return err
			}

			for _, shelf := range shelvestates {
				observer.ObserveInt64(m.overallFan, ternary(shelf.overallFan), metric.WithAttributes(shelf.attributes...))
				observer.ObserveInt64(m.overallTemp, ternary(shelf.overallTemp), metric.WithAttributes(shelf.attributes...))
				observer.ObserveInt64(m.overallPSU, ternary(shelf.overallPSU), metric.WithAttributes(shelf.attributes...))

				for _, sensor := range shelf.sensors {
					attributes := make([]attribute.KeyValue, 0, len(shelf.attributes)+len(sensor.attributes()))
					attributes = append(attributes, shelf.attributes...)
					attributes = append(attributes, sensor.attributes()...)
					observer.ObserveInt64(m.sensorStatus, ternary(sensor.ok()), metric.WithAttributes(attributes...))
					observer.ObserveInt64(m.sensorValue, sensor.sensor.Value, metric.WithAttributes(attributes...))
				}
			}

			return nil
		},
		m.overallPSU, m.overallTemp, m.overallFan, m.sensorStatus, m.sensorValue,
	)
	if err != nil {
		m.logger.Error("error registering shelf metrics", slog.String("error", err.Error()))
	}

	return err
}

func (m *ShelfMetrics) updateShelfLocations(ctx context.Context) error {
	if len(m.shelfLocations.shelfLocations) > 0 {
		// precheck to see if its already been done
		return nil
	}

	m.shelfLocations.shelfLocationsLock.Lock()
	defer m.shelfLocations.shelfLocationsLock.Unlock()

	if len(m.shelfLocations.shelfLocations) > 0 {
		// already done by another first call to observe metrics
		return nil
	}

	disks, err := m.client.DiskService().GetDisks(ctx)
	if err != nil {
		return err
	}

	m.shelfLocations.shelfLocations = map[string]string{}

	for _, disk := range disks {
		// Remove the Controller designation from the ShelfLocation so we dont get duplicate
		// series if the controllers active state changes.
		shelfLocation := disk.ShelfLocation

		shelfLoationParts := strings.SplitN(disk.ShelfLocation, ".", 2)
		if len(shelfLoationParts) > 1 {
			shelfLocation = shelfLoationParts[1]
		}

		m.shelfLocations.shelfLocations[disk.ShelfID] = shelfLocation
	}

	return nil
}

func (m *ShelfMetrics) getShelfLocation(id string) string {
	m.shelfLocations.shelfLocationsLock.RLock()
	defer m.shelfLocations.shelfLocationsLock.RUnlock()

	return m.shelfLocations.shelfLocations[id]
}

type shelfSensor struct {
	sensor models.ShelfSensor
}

func (s *shelfSensor) ok() bool {
	return s.sensor.Status == "OK"
}

func (s *shelfSensor) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		{Key: "chassis_id", Value: attribute.StringValue(s.sensor.CID)},
		{Key: "name", Value: attribute.StringValue(s.sensor.Name)},
		{Key: "location", Value: attribute.StringValue(s.sensor.Location)},
		{Key: "type", Value: attribute.StringValue(s.sensor.Type)},
	}
}

type shelvestate struct {
	overallPSU  bool
	overallFan  bool
	overallTemp bool
	sensors     []shelfSensor

	attributes []attribute.KeyValue
}

func (m *ShelfMetrics) getShelfStates(ctx context.Context) ([]shelvestate, error) {
	service := m.client.ShelfService()

	shelves, err := service.GetShelves(ctx)
	if err != nil {
		return nil, err
	}

	shelvestates := make([]shelvestate, 0, len(shelves))

	for _, shelf := range shelves {
		shelfLocationId := m.getShelfLocation(shelf.ID)
		attributes := []attribute.KeyValue{
			{Key: "model", Value: attribute.StringValue(shelf.Model)},
			{Key: "serial", Value: attribute.StringValue(shelf.Serial)},
			{Key: "type", Value: attribute.StringValue(shelf.ChassisType)},
		}

		if shelfLocationId != "" {
			attributes = append(
				attributes,
				attribute.KeyValue{Key: "shelf", Value: attribute.StringValue(shelfLocationId)},
			)
		}

		sensors := []shelfSensor{}

		for _, sensor := range shelf.ChassisSensors {
			sensors = append(sensors, shelfSensor{sensor: sensor})
		}

		for _, ctrlr := range shelf.Ctrlrs {
			for _, sensor := range ctrlr.Sensors {
				sensors = append(sensors, shelfSensor{sensor: sensor})
			}
		}

		shelvestates = append(shelvestates, shelvestate{
			overallPSU:  shelf.PSUOverallStatus == "OK",
			overallFan:  shelf.FanOverallStatus == "OK",
			overallTemp: shelf.TempOverallStatus == "OK",

			sensors: sensors,

			attributes: attributes,
		})
	}

	return shelvestates, nil
}
