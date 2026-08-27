package exporter

import (
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/metric"

	"github.com/ChrisPortman/nimble_metric_exporter/pkg/client"
)

type NimbleCientFactory func(target string, logger *slog.Logger) (*client.NimbleClient, error)

type targetMetricPipeline struct {
	nimble      *client.NimbleClient
	meter       metric.Meter
	promHandler http.Handler
	logger      *slog.Logger
}

type PrometheusServerOpts struct {
	DisableDiskMetrics   bool
	DisableShelfMetrics  bool
	DisablePoolMetrics   bool
	DisableVolumeMetrics bool

	Logger *slog.Logger
}

type PrometheusServer struct {
	pipelines     map[string]*targetMetricPipeline
	lock          sync.RWMutex
	clientFactory NimbleCientFactory
	options       PrometheusServerOpts
	logger        *slog.Logger
}

func NewPrometheusServer(clientFactory NimbleCientFactory, opts PrometheusServerOpts) *PrometheusServer {
	logger := opts.Logger
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}

	return &PrometheusServer{
		pipelines:     make(map[string]*targetMetricPipeline),
		lock:          sync.RWMutex{},
		clientFactory: clientFactory,
		options:       opts,
		logger:        logger,
	}
}

func (s *PrometheusServer) Serve(listen string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.Handle("/", s)

	server := http.Server{
		Addr:              listen,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 5,
	}

	return server.ListenAndServe()
}

func (p *PrometheusServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	if target == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("'target' query option required"))

		return
	}

	metrics, err := p.targetPipeline(target)
	if err != nil {
		p.logger.Error(
			"unable to acquire metric pipeline",
			slog.String("target", target),
			slog.String("error", err.Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error: " + err.Error()))

		return
	}

	metrics.promHandler.ServeHTTP(w, r)
}

func (p *PrometheusServer) targetPipeline(target string) (*targetMetricPipeline, error) {
	if pipeline, exists := p.getPipeline(target); exists {
		return pipeline, nil
	}

	return p.createPipeline(target)
}

func (p *PrometheusServer) getPipeline(target string) (*targetMetricPipeline, bool) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	pipeline, exists := p.pipelines[target]

	if exists {
		pipeline.logger.Info("existing pipeline acquired")
	}

	return pipeline, exists
}

func (p *PrometheusServer) createPipeline(target string) (*targetMetricPipeline, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	targetLogger := p.logger.With(slog.String("target", target))

	// There is the potential that between calling get() with no result and
	// calling create(), that another operation created
	if pipeline, exists := p.pipelines[target]; exists {
		return pipeline, nil
	}

	targetLogger.Info("creating metric pipeline")

	registry := prometheus.NewRegistry()

	promExporter, err := NewPrometheusExporter(registry)
	if err != nil {
		return nil, err
	}

	provider, err := NewMetricsProvider(promExporter)
	if err != nil {
		return nil, err
	}

	meter := provider.Meter("github.com/ChrisPortman/nimble_metric_exporter/pkg/exporter")

	nimbleClient, err := p.clientFactory(target, targetLogger)
	if err != nil {
		return nil, err
	}

	pipeline := targetMetricPipeline{
		nimble:      nimbleClient,
		meter:       meter,
		promHandler: promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
		logger:      targetLogger,
	}

	if err := p.registerMetrics(&pipeline); err != nil {
		return nil, err
	}

	p.pipelines[target] = &pipeline

	return &pipeline, nil
}

func (p *PrometheusServer) registerMetrics(target *targetMetricPipeline) error {
	if !p.options.DisableDiskMetrics {
		target.logger.Info("registering disk metrics")

		_, err := NewDiskMetrics(target.nimble, target.meter, target.logger)
		if err != nil {
			return err
		}
	}

	if !p.options.DisableShelfMetrics {
		target.logger.Info("registering shelf metrics")

		_, err := NewShelfMetrics(target.nimble, target.meter, target.logger)
		if err != nil {
			return err
		}
	}

	if !p.options.DisablePoolMetrics {
		target.logger.Info("registering pool metrics")

		_, err := NewPoolMetrics(target.nimble, target.meter, target.logger)
		if err != nil {
			return err
		}
	}

	if !p.options.DisableVolumeMetrics {
		target.logger.Info("registering volume metrics")

		_, err := NewVolumeMetrics(target.nimble, target.meter, target.logger)
		if err != nil {
			return err
		}
	}

	return nil
}
