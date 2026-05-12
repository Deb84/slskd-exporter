package slskd

import (
	"fmt"
	"io"
	"net/http"
)

func (client *Client) NewRequest(method string, route string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, route, body)
	if err != nil {
		return nil, fmt.Errorf(`Unable to create a request. Route="%s": %w`, route, err)
	}

	req.Header = client.Headers

	return req, nil
}
