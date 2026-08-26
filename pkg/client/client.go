package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/ChrisPortman/nimble_metric_exporter/internal/models"
)

const nimbleAPIPort uint16 = 5392

type NimbleClient struct {
	Host string
	Port uint16

	username   string
	password   string
	version    string
	token      *models.Token
	httpClient http.Client
	logger     *slog.Logger
}

type NimbleClientOption func(c *NimbleClient)

func SetNimbleHost(h string) NimbleClientOption {
	return func(c *NimbleClient) {
		c.Host = h
	}
}

func SetNimblePort(p uint16) NimbleClientOption {
	return func(c *NimbleClient) {
		c.Port = p
	}
}

func SetTlsSkipVerify() NimbleClientOption {
	return func(c *NimbleClient) {
		c.httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, //#nosec
			},
		}
	}
}

func SetLogger(l *slog.Logger) NimbleClientOption {
	return func(c *NimbleClient) {
		c.logger = l
	}
}

func NewNimbleCLient(
	ctx context.Context,
	username string,
	password string,
	opts ...NimbleClientOption,
) (*NimbleClient, error) {
	client := NimbleClient{
		Host: "localhost",
		Port: nimbleAPIPort,

		username:   username,
		password:   password,
		httpClient: http.Client{},
		logger:     slog.New(slog.DiscardHandler),
	}

	for _, opt := range opts {
		opt(&client)
	}

	if err := client.versions(ctx); err != nil {
		return nil, err
	}

	if err := client.authenticate(ctx); err != nil {
		return nil, err
	}

	go client.authRefresher(ctx)

	return &client, nil
}

func (c *NimbleClient) DiskService() *DiskService {
	return &DiskService{client: c}
}

func (c *NimbleClient) ShelfService() *ShelfService {
	return &ShelfService{client: c}
}

func (c *NimbleClient) PoolService() *PoolService {
	return &PoolService{client: c}
}

func (c *NimbleClient) VolumeService() *VolumeService {
	return &VolumeService{client: c}
}

const versionPath string = "/versions"

func (c *NimbleClient) versions(ctx context.Context) error {
	base := url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   versionPath,
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), nil)
	if err != nil {
		return err
	}

	versionsData, err := doRequest[[]models.Version](c, request)
	if err != nil {
		return err
	}

	for _, v := range versionsData {
		c.version = v.Name

		break
	}

	if c.version == "" {
		return errors.New("unable to determine API version")
	}

	c.logger.Info("Storage API version", "version", c.version)

	return nil
}

func (c *NimbleClient) makeURL(path string) (url.URL, error) {
	fullPath, err := url.JoinPath(c.version, path)
	if err != nil {
		return url.URL{}, err
	}

	return url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   fullPath,
	}, err
}

const tokenPath string = "tokens"

func (c *NimbleClient) authenticate(ctx context.Context) error {
	tokenURL, err := c.makeURL(tokenPath)
	if err != nil {
		return err
	}

	requestData := models.APIRequest[models.CreateTokenRequest]{
		Data: models.CreateTokenRequest{
			Username: c.username,
			Password: c.password,
		},
	}

	request, err := newJsonHttpRequest(ctx, http.MethodPost, tokenURL, requestData)
	if err != nil {
		return err
	}

	token, err := doRequest[models.Token](c, request)
	if err != nil {
		return err
	}

	if token.SessionToken == "" {
		return errors.New("session token not provided")
	}

	c.token = &token

	return nil
}

func (c *NimbleClient) authRefresher(ctx context.Context) {
	for {
		expireTime := time.Now().Add(time.Minute)
		expireDuration := time.Duration(time.Until(expireTime).Seconds())

		if c.token != nil {
			expireTime = time.Unix(c.token.ExpiryTime, 0)
			expireDuration = time.Duration(float64(time.Second) * time.Until(expireTime).Seconds() * 0.9)
		}

		expireTimer := time.NewTimer(expireDuration)
		c.logger.Info("will refresh auth in in " + expireDuration.String())

		select {
		case <-expireTimer.C:
			c.logger.Debug("renewing authentication token")

			if err := c.authenticate(ctx); err != nil {
				c.logger.Error("error renewing token", slog.String("error", err.Error()))
				c.token = nil
			}
		case <-ctx.Done():
			c.logger.Debug("authentication refresher finished")

			return
		}
	}
}

func newJsonHttpRequest[T any](ctx context.Context, method string, dest url.URL, data T) (*http.Request, error) {
	jsonBytes, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	jsonReader := bytes.NewBuffer(jsonBytes)

	request, err := http.NewRequestWithContext(ctx, method, dest.String(), jsonReader)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return request, err
}

const tokenHeader string = "X-Auth-Token" //#nosec

func doRequest[T any](client *NimbleClient, req *http.Request) (T, error) {
	var (
		zero T
		data models.APIResponse[T]
	)

	if client.token != nil {
		req.Header.Set(tokenHeader, client.token.SessionToken)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return zero, err
	}

	defer func() { _ = resp.Body.Close() }()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return zero, err
	}

	if data.Error != nil {
		return zero, errors.New(data.Error.Text)
	}

	return data.Data, nil
}
