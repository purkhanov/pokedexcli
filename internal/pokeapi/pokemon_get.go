package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	var pokemonResp Pokemon
	var err error

	data, ok := c.cache.Get(url)
	if !ok {
		data, err = c.makeRequest(url)
		if err != nil {
			return pokemonResp, err
		}
	}

	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return pokemonResp, err
	}

	c.cache.Add(url, data)

	return pokemonResp, nil
}
