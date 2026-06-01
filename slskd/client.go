package slskd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type Client struct {
	HttpClient *http.Client
	Routes     *Routes
	Headers    http.Header
}

func handleCert(certPath string) (*tls.Config, error) {
	if certPath == "" {
		slog.Warn("No certificate provided, will ignore certificate verification")
		return &tls.Config{
			InsecureSkipVerify: true,
		}, nil
	}

	certPool := x509.NewCertPool()

	pemCert, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	certPool.AppendCertsFromPEM(pemCert)

	return &tls.Config{
		RootCAs: certPool,
	}, nil
}

func NewClient(routes *Routes, certPath string) *Client {

	tlsConfig, err := handleCert(certPath)
	if err != nil {
		slog.Warn(fmt.Sprintf("Invalid Certificate: %s Error: %s", certPath, err))
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return &Client{
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
