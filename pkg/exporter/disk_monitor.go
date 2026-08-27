package exporter

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ChrisPortman/nimble_metric_exporter/pkg/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type DiskMetrics struct {
	service *client.DiskService
	logger  *slog.Logger

	failed metric.Int64ObservableGauge
	absent metric.Int64ObservableGauge
	raid   metric.Int64ObservableGauge
	size   metric.Int64ObservableGauge
}

func NewDiskMetrics(service *client.DiskService, meter metric.Meter, logger *slog.Logger) (DiskMetrics, error) {
	var err error

	log := slog.New(slog.DiscardHandler)
	if logger != nil {
		log = logger
	}

	metrics := DiskMetrics{
		service: service,
		logger:  log,
	}

	metrics.absent, err = meter.Int64ObservableGauge(
		"nimble.disk.state.absent",
		metric.WithDescription("Indicates disk is in failed state."),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.disk.state.absent: %w", err)
	}

	metrics.failed, err = meter.Int64ObservableGauge(
		"nimble.disk.state.failed",
		metric.WithDescription("Indicates disk is in failed state."),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.disk.state.failed: %w", err)
	}

	metrics.raid, err = meter.Int64ObservableGauge(
		"nimble.disk.raid.resync",
		metric.WithDescription("Indicates disk is in raid resyncing."),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.disk.raid.resync: %w", err)
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

			diskStates, err := getDiskStates(ctx, m.service)
			if err != nil {
				m.logger.Error("error retrieving disk data", slog.String("error", err.Error()))

				return err
			}

			for _, disk := range diskStates {
				observer.ObserveInt64(m.absent, disk.absent, metric.WithAttributes(disk.attributes...))
				observer.ObserveInt64(m.failed, disk.failed, metric.WithAttributes(disk.attributes...))
				observer.ObserveInt64(m.raid, disk.raid, metric.WithAttributes(disk.attributes...))
				observer.ObserveInt64(m.size, disk.size, metric.WithAttributes(disk.attributes...))
			}

			return nil
		},
		m.absent, m.failed, m.raid, m.size,
	)
	if err != nil {
		m.logger.Error("error registering disk metrics", slog.String("error", err.Error()))
	}

	return err
}

type diskState struct {
	absent int64
	failed int64
	raid   int64
	size   int64

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
			failed        int64
			absent        int64
			raidResyncing int64
		)

		if disk.State == "failed" || disk.State == "t_fail" {
			failed = 1
		}

		if disk.State == "absent" || disk.State == "removed" {
			absent = 1
		}

		if disk.RaidState == "resynchronizing" {
			raidResyncing = 1
		}

		attributes := []attribute.KeyValue{
			{Key: "model", Value: attribute.StringValue(disk.Model)},
			{Key: "serial", Value: attribute.StringValue(disk.Serial)},
			{Key: "shelf", Value: attribute.Int64Value(disk.ShelfLocationID)},
			{Key: "slot", Value: attribute.Int64Value(disk.Slot)},
			{Key: "type", Value: attribute.StringValue(disk.Type)},
		}

		diskStates = append(diskStates, diskState{
			absent:     absent,
			failed:     failed,
			raid:       raidResyncing,
			size:       disk.Size,
			attributes: attributes,
		})
	}

	return diskStates, nil
}
