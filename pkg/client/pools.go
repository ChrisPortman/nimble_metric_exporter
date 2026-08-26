package client

import (
	"context"
	"net/http"

	"github.com/ChrisPortman/nimble_metric_exporter/internal/models"
)

type PoolService struct {
	client *NimbleClient
}

const poolsPath string = "pools/detail"

func (c *PoolService) GetPools(ctx context.Context) ([]models.Pool, error) {
	poolsURL, err := c.client.makeURL(poolsPath)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, poolsURL.String(), nil)
	if err != nil {
		return nil, err
	}

	pools, err := doRequest[[]models.Pool](c.client, request)
	if err != nil {
		return nil, err
	}

	return pools, nil
}
