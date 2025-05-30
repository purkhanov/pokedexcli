package main

import (
	"time"

	"github.com/purkhanov/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Millisecond*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
