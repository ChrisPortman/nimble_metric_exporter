package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.43.0"
)

// NewMetricsPipeline creates a new independent metrics pipeline.
func NewMetricsProvider(exporter metric.Reader) (*metric.MeterProvider, error) {
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("nimble-metrics-exporter"),
			semconv.ServiceVersion("0.1.0"),
		),
	)
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(exporter),
	)

	return meterProvider, nil
}

// NewPrometheusExporter returns a new otelprom.Exporter that is a Prometheus exporter
// backed by a dedicated prometheus registry
func NewPrometheusExporter(registry *prometheus.Registry) (*otelprom.Exporter, error) {
	return otelprom.New(
		otelprom.WithRegisterer(registry),
		otelprom.WithoutScopeInfo(),
	)
}
