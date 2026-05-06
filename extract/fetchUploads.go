package extract

import (
	"encoding/json"
	"fmt"
	"slskd-exporter/domain/slskd"
)

func FetchUploads(client *Client) {

	req, err := client.NewRequest("GET", client.Routes.Uploads, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.HttpClient.Do(req)
	if err != nil {
		panic(err)
	}

	//body, err := io.ReadAll(response.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(body))

	var transfers slskd.Transfers
	json.NewDecoder(response.Body).Decode(&transfers)
	fmt.Println(transfers)
}
