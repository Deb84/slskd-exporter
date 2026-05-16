package slskd

import (
	"crypto/tls"
	"net/http"
	"slskd-exporter/logger"
)

type Client struct {
	Logger     *logger.Logger
	HttpClient *http.Client
	Routes     *Routes
	Headers    http.Header
}

func NewClient(logger *logger.Logger, routes *Routes) *Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return &Client{
		Logger:     logger,
		HttpClient: httpClient,
		Routes:     routes,
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
}

func (client *Client) SetHeader(key string, value []string) {
	client.Headers[key] = value
}
