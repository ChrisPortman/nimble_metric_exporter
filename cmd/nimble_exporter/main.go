// This go:debug needs to be set. At least at last check (2026-08-26) an HPE Alletra 5000
// required this to be set.

//go:debug tlsrsakex=1
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/alecthomas/kong"

	"github.com/ChrisPortman/nimble_metric_exporter/internal/version"
	"github.com/ChrisPortman/nimble_metric_exporter/pkg/client"
	"github.com/ChrisPortman/nimble_metric_exporter/pkg/exporter"
)

type CLI struct {
	Host        string `name:"host" short:"h" env:"NIMBLE_HOST" help:"the IP or hostname of the Nimble array"`
	Username    string `name:"username" short:"u" env:"NIMBLE_USERNAME" help:"Username to use when accessing the array"`
	Password    string `name:"password" short:"p" env:"NIMBLE_PASSWORD" help:"Password to use when accessing the array"`
	SSLNoVerify bool   `name:"ssl-no-verify" env:"NIMBLE_SSL_NO_VERIFY" help:"Disable TLS certificate verification when connecting to the array"`
	Listen      string `name:"listen" short:"l" env:"NIMBLE_LISTEN" help:"Set the prometheus scrape interface to listen on IP and port" default:":9490"`

	ExcludeShelf  bool `name:"exclude-shelf" env:"NIMBLE_EXCLUDE_SHELF" help:"Disable collection of shelf hardware metrics"`
	ExcludeDisk   bool `name:"exclude-disk" env:"NIMBLE_EXCLUDE_DISK" help:"Disable collection of disk metrics"`
	ExcludePool   bool `name:"exclude-pool" env:"NIMBLE_EXCLUDE_POOL" help:"Disable collection of pool metrics"`
	ExcludeVolume bool `name:"exclude-volume" env:"NIMBLE_EXCLUDE_VOLUME" help:"Disable collection of volume metrics"`

	Debug bool `name:"debug" short:"d" help:"Enable debug mode."`
}

func (c *CLI) Run() error {
	ctx, cnl := context.WithCancel(context.Background())
	defer cnl()

	var logLevel slog.LevelVar

	if c.Debug {
		logLevel.Set(slog.LevelDebug)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: &logLevel}))

	logger.Info(
		"starting Nimble Metric Exporter",
		slog.String("version", version.Version),
		slog.String("build_date", version.BuildDate),
		slog.String("go_version", version.GoVersion),
	)

	otelClose, err := exporter.InitaliseOtel()
	if err != nil {
		return err
	}
	defer otelClose()

	clientOpts := []client.NimbleClientOption{
		client.SetNimbleHost(c.Host),
		client.SetLogger(logger),
	}
	if c.SSLNoVerify {
		clientOpts = append(clientOpts, client.SetTlsSkipVerify())
	}

	client, err := client.NewNimbleCLient(
		ctx, c.Username, c.Password, clientOpts...,
	)
	if err != nil {
		return err
	}

	if err := c.registerMetrics(ctx, client, logger); err != nil {
		return err
	}

	logger.Info("starting prometheus scrape endpoint", slog.String("address", c.Listen))

	if err := exporter.ServePrometheus(c.Listen); err != nil {
		logger.Error(err.Error())
	}

	return err
}

func (c *CLI) registerMetrics(ctx context.Context, client *client.NimbleClient, logger *slog.Logger) error {
	if !c.ExcludeDisk {
		logger.Info("registering disk metrics")

		diskMetrics, err := exporter.NewDiskMetrics(client.DiskService(), logger)
		if err != nil {
			return err
		}

		if err := diskMetrics.Register(ctx); err != nil {
			return err
		}
	}

	if !c.ExcludeShelf {
		logger.Info("registering shelf metrics")

		shelfMetrics, err := exporter.NewShelfMetrics(client.ShelfService(), logger)
		if err != nil {
			return err
		}

		if err := shelfMetrics.Register(ctx); err != nil {
			return err
		}
	}

	if !c.ExcludePool {
		logger.Info("registering pool metrics")

		poolMetrics, err := exporter.NewPoolMetrics(client.PoolService(), logger)
		if err != nil {
			return err
		}

		if err := poolMetrics.Register(ctx); err != nil {
			return err
		}
	}

	if !c.ExcludeVolume {
		logger.Info("registering volume metrics")

		volumeMetrics, err := exporter.NewVolumeMetrics(client.VolumeService(), logger)
		if err != nil {
			return err
		}

		if err := volumeMetrics.Register(ctx); err != nil {
			return err
		}
	}

	return nil
}

var cli CLI

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
