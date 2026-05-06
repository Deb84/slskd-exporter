package extract

import (
	"crypto/tls"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	Routes     *Routes
	Headers    http.Header
}

func NewClient(routes *Routes) Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return Client{
		HttpClient: httpClient,
		Routes:     routes,
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
}

func (c *Client) SetHeader(key string, value []string) {
	c.Headers[key] = value
}
