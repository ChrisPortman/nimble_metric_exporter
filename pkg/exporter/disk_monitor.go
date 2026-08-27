package exporter

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ChrisPortman/nimble_metric_exporter/pkg/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type DiskMetrics struct {
	client *client.NimbleClient
	logger *slog.Logger

	stateOK metric.Int64ObservableGauge
	raidOK  metric.Int64ObservableGauge
	size    metric.Int64ObservableGauge
}

func NewDiskMetrics(client *client.NimbleClient, meter metric.Meter, logger *slog.Logger) (DiskMetrics, error) {
	var err error

	log := slog.New(slog.DiscardHandler)
	if logger != nil {
		log = logger
	}

	metrics := DiskMetrics{
		client: client,
		logger: log,
	}

	metrics.stateOK, err = meter.Int64ObservableGauge(
		"nimble.disk.state.ok",
		metric.WithDescription("Indicates disk is in failed state when value is 0."),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.disk.state.ok: %w", err)
	}

	metrics.raidOK, err = meter.Int64ObservableGauge(
		"nimble.disk.raid.ok",
		metric.WithDescription("Indicates disk is in raid state issus when value is 0."),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.disk.raid.ok: %w", err)
	}

	metrics.size, err = meter.Int64ObservableGauge(
		"nimble.disk.size",
		metric.WithDescription("Disk size in bytes."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.disk.size: %w", err)
	}

	if err := metrics.Register(meter); err != nil {
		return metrics, err
	}

	return metrics, nil
}

func (m *DiskMetrics) Register(meter metric.Meter) error {
	_, err := meter.RegisterCallback(
		func(ctx context.Context, observer metric.Observer) error {
			m.logger.Debug("loading disk metrics")

			diskStates, err := getDiskStates(ctx, m.client.DiskService())
			if err != nil {
				m.logger.Error("error retrieving disk data", slog.String("error", err.Error()))

				return err
			}

			for _, disk := range diskStates {
				observer.ObserveInt64(m.stateOK, disk.stateOK, metric.WithAttributes(disk.attributes...))
				observer.ObserveInt64(m.size, disk.size, metric.WithAttributes(disk.attributes...))

				if disk.raidValid {
					observer.ObserveInt64(m.raidOK, disk.raidOK, metric.WithAttributes(disk.attributes...))
				}
			}

			return nil
		},
		m.stateOK, m.raidOK, m.size,
	)
	if err != nil {
		m.logger.Error("error registering disk metrics", slog.String("error", err.Error()))
	}

	return err
}

type diskState struct {
	stateOK   int64
	raidOK    int64
	raidValid bool
	size      int64

	attributes []attribute.KeyValue
}

func getDiskStates(ctx context.Context, service *client.DiskService) ([]diskState, error) {
	disks, err := service.GetDisks(ctx)
	if err != nil {
		return nil, err
	}

	diskStates := make([]diskState, 0, len(disks))

	for _, disk := range disks {
		var (
			stateOK   int64
			raidOK    int64
			raidValid bool
		)

		if disk.State != "failed" && disk.State != "t_fail" && disk.State != "absent" {
			stateOK = 1
		}

		if disk.RaidID >= 0 {
			raidValid = true

			if disk.RaidState == "okay" {
				raidOK = 1
			}
		}

		// Remove the Controller designation from the ShelfLocation so we dont get duplicate
		// series if the controllers active state changes.
		shelfLocation := disk.ShelfLocation

		shelfLoationParts := strings.SplitN(shelfLocation, ".", 2)
		if len(shelfLoationParts) > 1 {
			shelfLocation = shelfLoationParts[1]
		}

		attributes := []attribute.KeyValue{
			{Key: "model", Value: attribute.StringValue(disk.Model)},
			{Key: "serial", Value: attribute.StringValue(disk.Serial)},
			{Key: "shelf", Value: attribute.StringValue(shelfLocation)},
			{Key: "slot", Value: attribute.Int64Value(disk.Slot)},
			{Key: "type", Value: attribute.StringValue(disk.Type)},
		}

		diskStates = append(diskStates, diskState{
			stateOK:    stateOK,
			raidOK:     raidOK,
			raidValid:  raidValid,
			size:       disk.Size,
			attributes: attributes,
		})
	}

	return diskStates, nil
}
