package extract

import (
	"io"
	"net/http"
)

func (client *Client) NewRequest(method string, route string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, route, body)
	if err != nil {
		panic(err)
	}

	req.Header = client.Headers

	return req, err
}
