package slskd

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"slskd-exporter/models/slskd"
	"time"
)

type Session struct {
	Username      string
	Credentials   slskd.Credentials
	Authorization slskd.Authorization
}

func NewSession(client *Client, user string, pass string) *Session {
	payload := slskd.Credentials{
		Password: pass,
		Username: user,
	}

	session := &Session{
		Username:    user,
		Credentials: payload,
	}

	err := session.NewToken(client, &payload)
	if err != nil {
		slog.Error("Unable to get a Slskd token")
	}

	return session
}

func (session *Session) NewToken(client *Client, payload *slskd.Credentials) error {
	routes := client.Routes

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := client.NewRequest("POST", routes.Session, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	//fmt.Println(response.StatusCode)
	var token slskd.Authorization
	json.NewDecoder(response.Body).Decode(&token)
	session.Authorization = token

	return err
}

func (session *Session) IsExpiredToken(margin time.Duration) bool {
	expires, _ := time.Parse(time.RFC3339, session.Authorization.Expires)
	return time.Now().After(expires.Add(margin))
}

func (session *Session) RenewToken(client *Client) {
	if session.NewToken(client, &session.Credentials) != nil {
		slog.Error("Unable to renew the Slskd token")
	}
}
