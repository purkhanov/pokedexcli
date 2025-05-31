package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationsResp RespShallowLocations
	var err error

	data, ok := c.cache.Get(url)
	if !ok {
		data, err = c.makeRequest(url)
		if err != nil {
			return locationsResp, err
		}
	}

	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
