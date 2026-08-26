package client

import (
	"context"
	"net/http"

	"github.com/ChrisPortman/nimble_metric_exporter/internal/models"
)

type ShelfService struct {
	client *NimbleClient
}

const shelvesPath string = "shelves/detail"

func (c *ShelfService) GetShelves(ctx context.Context) ([]models.Shelf, error) {
	disksURL, err := c.client.makeURL(shelvesPath)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, disksURL.String(), nil)
	if err != nil {
		return nil, err
	}

	disks, err := doRequest[[]models.Shelf](c.client, request)
	if err != nil {
		return nil, err
	}

	return disks, nil
}
