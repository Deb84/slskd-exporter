package extract

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"slskd-exporter/domain"
)

func Extract(env domain.SlskdEnv) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	routes := CreateRoutes(env.HOST, env.PORT)
	token := Auth(*httpClient, routes, env.USER, env.PASSWORD)
	fmt.Println(token)
}
