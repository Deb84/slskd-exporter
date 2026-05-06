package extract

import (
	"slskd-exporter/domain"
)

func Extract(env domain.SlskdEnv) {
	routes := CreateRoutes(env.HOST, env.PORT)
	client := NewClient(&routes)

	token := Auth(&client, env.USER, env.PASSWORD)
	client.SetHeader("Authorization", []string{token.TokenType + " " + token.Token})
	FetchUploads(&client)
}
