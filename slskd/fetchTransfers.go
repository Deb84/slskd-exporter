package slskd

import (
	"encoding/json"
	"fmt"
	"slskd-exporter/models/slskd"
)

func FetchTransfer(client *Client, route string) (*slskd.Transfers, error) {
	req, err := client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(`Unable to reach "%s": %w`, route, err)
	}

	var transfers slskd.Transfers
	json.NewDecoder(response.Body).Decode(&transfers)

	return &transfers, nil
}

func FetchTransfers(client *Client) (*slskd.ExtractedTransfers, error) {
	uploads, err := FetchTransfer(client, client.Routes.Uploads)
	if err != nil {
		return nil, err
	}
	downloads, err := FetchTransfer(client, client.Routes.Downloads)
	if err != nil {
		return nil, err
	}

	return &slskd.ExtractedTransfers{
		Downloads: downloads,
		Uploads:   uploads,
	}, nil
}
