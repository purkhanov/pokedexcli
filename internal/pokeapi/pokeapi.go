package pokeapi

import (
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) makeRequest(urlPath string) ([]byte, error) {
	url := baseURL + urlPath

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return dat, nil
}
