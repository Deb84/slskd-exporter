package extract

import (
	"bytes"
	"encoding/json"
	"slskd-exporter/domain/slskd"
)

type authPayload struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Auth(client *Client, user string, pass string) slskd.Token {
	routes := client.Routes

	payload := authPayload{
		Password: pass,
		Username: user,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := client.NewRequest("POST", routes.Session, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	response, err := client.HttpClient.Do(req)
	if err != nil {
		panic(err)
	}

	//fmt.Println(response.StatusCode)
	var token slskd.Token
	json.NewDecoder(response.Body).Decode(&token)
	//fmt.Println(result.Token)

	defer response.Body.Close()

	return token
}
