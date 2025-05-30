package pokeapi

import (
	"net/http"
	"time"

	"github.com/purkhanov/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCahce(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
