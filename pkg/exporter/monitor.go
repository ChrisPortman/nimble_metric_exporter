package exporter

import (
	"context"
	"log/slog"
	"time"
)

type Monitor interface {
	Check(ctx context.Context) error
}

const intervalSecsDefault int = 15

type MonitorConfig struct {
	IntervalSecs int
	Logger       *slog.Logger
}

func RunMonitors(ctx context.Context, monitors []Monitor, cfg MonitorConfig) {
	interval := cfg.IntervalSecs
	if interval == 0 {
		interval = intervalSecsDefault
	}

	logger := slog.New(slog.DiscardHandler)
	if cfg.Logger != nil {
		logger = cfg.Logger
	}

	for {
		timer := time.NewTimer(time.Second * time.Duration(interval))
		select {
		case <-timer.C:
			for _, monitor := range monitors {
				go func() {
					if err := monitor.Check(ctx); err != nil {
						logger.Error(err.Error())
					}
				}()
			}
		case <-ctx.Done():
			return
		}
	}
}
