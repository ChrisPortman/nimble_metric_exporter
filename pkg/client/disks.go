package client

import (
	"context"
	"net/http"

	"github.com/ChrisPortman/nimble_metric_exporter/internal/models"
)

type DiskService struct {
	client *NimbleClient
}

const disksPath string = "disks/detail"

func (c *DiskService) GetDisks(ctx context.Context) ([]models.Disk, error) {
	disksURL, err := c.client.makeURL(disksPath)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, disksURL.String(), nil)
	if err != nil {
		return nil, err
	}

	disks, err := doRequest[[]models.Disk](c.client, request)
	if err != nil {
		return nil, err
	}

	return disks, nil
}
