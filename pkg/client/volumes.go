package client

import (
	"context"
	"net/http"

	"github.com/ChrisPortman/nimble_metric_exporter/internal/models"
)

type VolumeService struct {
	client *NimbleClient
}

const volumesPath string = "volumes/detail"

func (c *VolumeService) GetVolumes(ctx context.Context) ([]models.Volume, error) {
	volumesURL, err := c.client.makeURL(volumesPath)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, volumesURL.String(), nil)
	if err != nil {
		return nil, err
	}

	volumes, err := doRequest[[]models.Volume](c.client, request)
	if err != nil {
		return nil, err
	}

	return volumes, nil
}
