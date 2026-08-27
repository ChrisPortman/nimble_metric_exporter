package exporter

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ChrisPortman/nimble_metric_exporter/pkg/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type PoolMetrics struct {
	client *client.NimbleClient
	logger *slog.Logger

	capacity      metric.Int64ObservableGauge
	used          metric.Int64ObservableGauge
	volumes       metric.Int64ObservableGauge
	cacheCapacity metric.Int64ObservableGauge
}

func NewPoolMetrics(client *client.NimbleClient, meter metric.Meter, logger *slog.Logger) (PoolMetrics, error) {
	var err error

	log := slog.New(slog.DiscardHandler)
	if logger != nil {
		log = logger
	}

	metrics := PoolMetrics{
		client: client,
		logger: log,
	}

	metrics.capacity, err = meter.Int64ObservableGauge(
		"nimble.pool.capacity",
		metric.WithDescription("Capacity of the pool in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.pool.capacity: %w", err)
	}

	metrics.used, err = meter.Int64ObservableGauge(
		"nimble.pool.used",
		metric.WithDescription("Usage of the pool in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.pool.used: %w", err)
	}

	metrics.volumes, err = meter.Int64ObservableGauge(
		"nimble.pool.volumes",
		metric.WithDescription("Number of volumes provisioned in the pool"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.pool.volumes: %w", err)
	}

	metrics.cacheCapacity, err = meter.Int64ObservableGauge(
		"nimble.pool.cache",
		metric.WithDescription("Capacity of cache in the pool"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return metrics, fmt.Errorf("error creating nimble.pool.cache: %w", err)
	}

	if err := metrics.Register(meter); err != nil {
		return metrics, err
	}

	return metrics, nil
}

func (m *PoolMetrics) Register(meter metric.Meter) error {
	_, err := meter.RegisterCallback(
		func(ctx context.Context, observer metric.Observer) error {
			m.logger.Debug("loading pool metrics")

			poolStates, err := getPoolStates(ctx, m.client.PoolService())
			if err != nil {
				m.logger.Error("error retrieving pool data", slog.String("error", err.Error()))

				return err
			}

			for _, pool := range poolStates {
				observer.ObserveInt64(m.capacity, pool.capacity, metric.WithAttributes(pool.attributes...))
				observer.ObserveInt64(m.used, pool.used, metric.WithAttributes(pool.attributes...))
				observer.ObserveInt64(m.volumes, pool.volumes, metric.WithAttributes(pool.attributes...))
				observer.ObserveInt64(m.cacheCapacity, pool.cacheCapacity, metric.WithAttributes(pool.attributes...))
			}

			return nil
		},
		m.capacity, m.used, m.volumes, m.cacheCapacity,
	)
	if err != nil {
		m.logger.Error("error registering disk metrics", slog.String("error", err.Error()))
	}

	return err
}

type poolState struct {
	capacity      int64
	used          int64
	volumes       int64
	cacheCapacity int64

	attributes []attribute.KeyValue
}

func getPoolStates(ctx context.Context, service *client.PoolService) ([]poolState, error) {
	pools, err := service.GetPools(ctx)
	if err != nil {
		return nil, err
	}

	poolStates := make([]poolState, 0, len(pools))

	for _, pool := range pools {
		attributes := []attribute.KeyValue{
			{Key: "name", Value: attribute.StringValue(pool.Name)},
		}

		// #nosec G115 uint64 is not supported by the OTEL framework
		poolStates = append(poolStates, poolState{
			capacity:      int64(pool.Capacity),
			used:          int64(pool.Usage),
			volumes:       int64(pool.VolCount),
			cacheCapacity: int64(pool.CacheCapacity),

			attributes: attributes,
		})
	}

	return poolStates, nil
}
