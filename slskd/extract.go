package slskd

import (
	"fmt"
	"slskd-exporter/logger"
	"slskd-exporter/models"
	"slskd-exporter/models/slskd"
	"slskd-exporter/retry"
	"time"
)

type ExtractionSession struct {
	Logger     *logger.Logger
	HttpClient *Client
	Session    *Session
}

func NewSessionWithRetry(logger *logger.Logger, client *Client, env *models.SlskdEnv) (*Session, error) {
	cb := func(count int) (*Session, error) {

		session, err := NewSession(client, env.USER, env.PASSWORD)
		if err == nil {
			return session, nil
		}

		logger.Warn("Unable to connect to slskd...", "Try", count)
		return nil, err
	}

	doneCB := func(err error, _ int) error {
		return err
	}

	session, err := retry.RetryWithTimeout(
		30*time.Second,
		5*time.Second,
		1,
		cb,
		doneCB,
	)

	if err != nil {
		return nil, err
	}
	return session, nil
}

func NewExtraction(logger *logger.Logger, env models.SlskdEnv) (*ExtractionSession, error) {
	routes := NewRoutes(env.HOST, env.PORT)
	client := NewClient(logger, &routes, env.CERT)
	session, err := NewSessionWithRetry(logger, client, &env)
	if err != nil {
		return nil, fmt.Errorf("Unable to create a new Slskd session: %w", err)
	}

	return &ExtractionSession{
		Logger:     logger,
		HttpClient: client,
		Session:    session,
	}, nil
}

func (extraction *ExtractionSession) DoExtraction() (*slskd.ExtractedTransfers, error) {
	client := extraction.HttpClient
	session := extraction.Session

	if session.IsExpiredToken(-5 * time.Minute) {
		err := session.RenewToken(client)
		if err != nil {
			extraction.Logger.Error("Renew token failed", "error", err)
		}
	}

	authorization := session.Authorization
	client.SetHeader("Authorization", []string{authorization.TokenType + " " + authorization.Token})

	transfers, err := FetchTransfers(client)
	if err != nil {
		return nil, fmt.Errorf("Slskd transfers extraction failed: %w", err)
	}

	return transfers, nil
}
