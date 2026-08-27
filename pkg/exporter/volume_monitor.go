package exporter

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ChrisPortman/nimble_metric_exporter/pkg/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type VolumeMetrics struct {
	client *client.NimbleClient
	logger *slog.Logger

	capacity    metric.Int64ObservableGauge
	used        metric.Int64ObservableGauge
	snapUsed    metric.Int64ObservableGauge
	connections metric.Int64ObservableGauge
	online      metric.Int64ObservableGauge

	encrypted metric.Int64ObservableGauge
	deduped   metric.Int64ObservableGauge

	readIops           metric.Int64ObservableGauge
	readThroughput     metric.Int64ObservableGauge
	readLatency        metric.Int64ObservableGauge
	writeIops          metric.Int64ObservableGauge
	writeThroughput    metric.Int64ObservableGauge
	writeLatency       metric.Int64ObservableGauge
	combinedIops       metric.Int64ObservableGauge
	combinedThroughput metric.Int64ObservableGauge
	combinedLatency    metric.Int64ObservableGauge
}

//nolint:cyclop,funlen
func NewVolumeMetrics(client *client.NimbleClient, meter metric.Meter, logger *slog.Logger) (VolumeMetrics, error) {
	var err error

	log := slog.New(slog.DiscardHandler)
	if logger != nil {
		log = logger
	}

	metrics := VolumeMetrics{
		client: client,
		logger: log,
	}

	metrics.capacity, err = meter.Int64ObservableGauge(
		"nimble.volume.capacity",
		metric.WithDescription("Capacity of the volume in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.capacity: %w", err)
	}

	metrics.used, err = meter.Int64ObservableGauge(
		"nimble.volume.used",
		metric.WithDescription("Usage of the volume in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.used: %w", err)
	}

	metrics.snapUsed, err = meter.Int64ObservableGauge(
		"nimble.volume.snapused",
		metric.WithDescription("Usage of the volume for snaps in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.snapused: %w", err)
	}

	metrics.connections, err = meter.Int64ObservableGauge(
		"nimble.volume.connections",
		metric.WithDescription("Number of connections to the volume"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.connections: %w", err)
	}

	metrics.online, err = meter.Int64ObservableGauge(
		"nimble.volume.online",
		metric.WithDescription("Indicates that the volume is online"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.online: %w", err)
	}

	metrics.encrypted, err = meter.Int64ObservableGauge(
		"nimble.volume.encrypted",
		metric.WithDescription("Indicates that the volume is encryption enabled"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.encrypted: %w", err)
	}

	metrics.deduped, err = meter.Int64ObservableGauge(
		"nimble.volume.deduped",
		metric.WithDescription("Indicates that the volume is deduplication enabled"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.deduped: %w", err)
	}

	metrics.readIops, err = meter.Int64ObservableGauge(
		"nimble.volume.read.iops.average.5m",
		metric.WithDescription("Average read IOPS over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.read.iops.average.5m: %w", err)
	}

	metrics.readThroughput, err = meter.Int64ObservableGauge(
		"nimble.volume.read.throughput.average.5m",
		metric.WithDescription("Average read throughput over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.read.throughput.average.5m: %w", err)
	}

	metrics.readLatency, err = meter.Int64ObservableGauge(
		"nimble.volume.read.latency.average.5m",
		metric.WithDescription("Average read latency over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.read.latency.average.5m: %w", err)
	}

	metrics.writeIops, err = meter.Int64ObservableGauge(
		"nimble.volume.write.iops.average.5m",
		metric.WithDescription("Average write IOPS over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.write.iops.average.5m: %w", err)
	}

	metrics.writeThroughput, err = meter.Int64ObservableGauge(
		"nimble.volume.write.throughput.average.5m",
		metric.WithDescription("Average write throughput over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.write.throughput.average.5m: %w", err)
	}

	metrics.writeLatency, err = meter.Int64ObservableGauge(
		"nimble.volume.write.latency.average.5m",
		metric.WithDescription("Average write latency over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.write.latency.average.5m: %w", err)
	}

	metrics.combinedIops, err = meter.Int64ObservableGauge(
		"nimble.volume.combined.iops.average.5m",
		metric.WithDescription("Average combined IOPS over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.combined.iops.average.5m: %w", err)
	}

	metrics.combinedThroughput, err = meter.Int64ObservableGauge(
		"nimble.volume.combined.throughput.average.5m",
		metric.WithDescription("Average combined throughput over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.combined.throughput.average.5m: %w", err)
	}

	metrics.combinedLatency, err = meter.Int64ObservableGauge(
		"nimble.volume.combined.latency.average.5m",
		metric.WithDescription("Average combined latency over last 5 minutes"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.volume.combined.latency.average.5m: %w", err)
	}

	if err := metrics.Register(meter); err != nil {
		return metrics, err
	}

	return metrics, nil
}

func (m *VolumeMetrics) Register(meter metric.Meter) error {
	ternary := func(in bool) int64 {
		if in {
			return 1
		}

		return 0
	}

	_, err := meter.RegisterCallback(
		func(ctx context.Context, observer metric.Observer) error {
			m.logger.Debug("loading volume metrics")

			volumeStates, err := getVolumeStates(ctx, m.client.VolumeService())
			if err != nil {
				m.logger.Error("error retrieving volume data", slog.String("error", err.Error()))

				return err
			}

			// #nosec G115 uint64 is not supported by the OTEL framework
			for _, volume := range volumeStates {
				observer.ObserveInt64(m.capacity, int64(volume.capacity), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.used, int64(volume.used), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.snapUsed, int64(volume.snapUsed), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.connections, int64(volume.connections), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.online, ternary(volume.online), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.encrypted, ternary(volume.encrypted), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.deduped, ternary(volume.deduped), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.readIops, int64(volume.readIops), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.readThroughput, int64(volume.readThroughput), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.readLatency, int64(volume.readLatency), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.writeIops, int64(volume.writeIops), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.writeThroughput, int64(volume.writeThroughput), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.writeLatency, int64(volume.writeLatency), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.combinedIops, int64(volume.combinedIops), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(m.combinedLatency, int64(volume.combinedLatency), metric.WithAttributes(volume.attributes...))
				observer.ObserveInt64(
					m.combinedThroughput,
					int64(volume.combinedThroughput),
					metric.WithAttributes(volume.attributes...),
				)
			}

			return nil
		},
		m.capacity,
		m.used,
		m.snapUsed,
		m.connections,
		m.online,
		m.encrypted,
		m.deduped,
		m.readIops,
		m.readThroughput,
		m.readLatency,
		m.writeIops,
		m.writeThroughput,
		m.writeLatency,
		m.combinedIops,
		m.combinedThroughput,
		m.combinedLatency,
	)
	if err != nil {
		m.logger.Error("error registering disk metrics", slog.String("error", err.Error()))
	}

	return err
}

type volumeState struct {
	capacity    uint64
	used        uint64
	snapUsed    uint64
	connections uint64
	online      bool

	encrypted bool
	deduped   bool

	readIops           uint64
	readThroughput     uint64
	readLatency        uint64
	writeIops          uint64
	writeThroughput    uint64
	writeLatency       uint64
	combinedIops       uint64
	combinedThroughput uint64
	combinedLatency    uint64

	attributes []attribute.KeyValue
}

func getVolumeStates(ctx context.Context, service *client.VolumeService) ([]volumeState, error) {
	volumes, err := service.GetVolumes(ctx)
	if err != nil {
		return nil, err
	}

	volumeStates := make([]volumeState, 0, len(volumes))

	for _, volume := range volumes {
		attributes := []attribute.KeyValue{
			{Key: "name", Value: attribute.StringValue(volume.Name)},
			{Key: "folder", Value: attribute.StringValue(volume.FolderName)},
			{Key: "pool", Value: attribute.StringValue(volume.PoolName)},
		}

		volumeStates = append(volumeStates, volumeState{
			capacity:    volume.Size,
			used:        volume.VolUsageUncompressedBytes,
			snapUsed:    volume.SnapUsageUncompressedBytes,
			connections: volume.NumConnections,
			online:      volume.Online,

			encrypted: volume.EncryptionCipher != "none",
			deduped:   volume.DedupeEnabled,

			readIops:           volume.AvgStatsLast5mins.ReadIops,
			readThroughput:     volume.AvgStatsLast5mins.ReadThroughput,
			readLatency:        volume.AvgStatsLast5mins.ReadLatency,
			writeIops:          volume.AvgStatsLast5mins.WriteIops,
			writeThroughput:    volume.AvgStatsLast5mins.WriteThroughput,
			writeLatency:       volume.AvgStatsLast5mins.WriteLatency,
			combinedIops:       volume.AvgStatsLast5mins.CombinedIops,
			combinedThroughput: volume.AvgStatsLast5mins.CombinedThroughput,
			combinedLatency:    volume.AvgStatsLast5mins.CombinedLatency,

			attributes: attributes,
		})
	}

	return volumeStates, nil
}
