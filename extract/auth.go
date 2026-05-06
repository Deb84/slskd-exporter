package extract

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type authPayload struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type authResponse struct {
	Token      string `json:"token"`
	TokenTyper string `json:"tokenType"`
}

func Auth(client http.Client, routes Routes, user string, pass string) string {

	payload := authPayload{
		Password: pass,
		Username: user,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", routes.Session, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//fmt.Println(response.StatusCode)
	var result authResponse
	json.NewDecoder(response.Body).Decode(&result)
	//fmt.Println(result.Token)

	defer response.Body.Close()

	return result.Token
}
