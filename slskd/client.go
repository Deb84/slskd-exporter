package slskd

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"slskd-exporter/logger"
)

type Client struct {
	Logger     *logger.Logger
	HttpClient *http.Client
	Routes     *Routes
	Headers    http.Header
}

func handleCert(logger *logger.Logger, certPath string) (*tls.Config, error) {
	if certPath == "" {
		logger.Warn("No certificate provided, will ignore certificate verification")
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

func NewClient(logger *logger.Logger, routes *Routes, certPath string) *Client {

	tlsConfig, err := handleCert(logger, certPath)
	if err != nil {
		logger.Warn("Invalid Certificate: %s Error: %s", certPath, err)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
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
