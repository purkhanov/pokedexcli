package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	var locationResp Location
	var err error

	data, ok := c.cache.Get(url)
	if !ok {
		data, err = c.makeRequest(url)
		if err != nil {
			return locationResp, err
		}
	}

	if err := json.Unmarshal(data, &locationResp); err != nil {
		return locationResp, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
