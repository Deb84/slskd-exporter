package slskd

import (
	"log/slog"
	"slskd-exporter/models"
	"slskd-exporter/models/slskd"
	"time"
)

type ExtractionSession struct {
	HttpClient *Client
	Session    *Session
}

func NewExtraction(env models.SlskdEnv) *ExtractionSession {
	routes := NewRoutes(env.HOST, env.PORT)
	client := NewClient(&routes, env.CERT)
	session := NewSession(client, env.USER, env.PASSWORD)

	return &ExtractionSession{
		HttpClient: client,
		Session:    session,
	}
}

func (extraction *ExtractionSession) DoExtraction() (*slskd.ExtractedTransfers, error) {
	client := extraction.HttpClient
	session := extraction.Session

	if session.IsExpiredToken(-5 * time.Minute) {
		session.RenewToken(client)
	}

	authorization := session.Authorization
	client.SetHeader("Authorization", []string{authorization.TokenType + " " + authorization.Token})

	transfers, err := FetchTransfers(client)
	if err != nil {
		slog.Error("An error occured with the slskd extraction.")
		return nil, err
	}

	return transfers, nil
}
