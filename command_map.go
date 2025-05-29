package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

// const baseUrl = "https://pokeapi.co/api/v2//location-area?offset=0&limit=20"

// var maps *pokeMap

// type pokeMap struct {
// 	Count    int       `json:"count"`
// 	Next     *string   `json:"next"`
// 	Previous *string   `json:"previous"`
// 	Results  []pokeRes `json:"results"`
// }

// type pokeRes struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// }

// func commandMap() error {
// 	url := baseUrl

// 	if maps != nil && maps.Next != nil {
// 		url = *maps.Next
// 	}

// 	pMap, err := makeRequest(url)
// 	if err != nil {
// 		return err
// 	}

// 	maps = &pMap

// 	printCities(pMap.Results)

// 	return nil
// }

// func commandMapb() error {
// 	if maps == nil || maps.Previous == nil {
// 		fmt.Println("you're on the first page")
// 		return nil
// 	}

// 	pMap, err := makeRequest(*maps.Previous)
// 	if err != nil {
// 		return err
// 	}

// 	maps = &pMap

// 	printCities(pMap.Results)

// 	return nil
// }

// func makeRequest(url string) (pokeMap, error) {
// 	var pokeMap pokeMap

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return pokeMap, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return pokeMap, err
// 	}

// 	if err := json.Unmarshal(body, &pokeMap); err != nil {
// 		return pokeMap, err
// 	}

// 	return pokeMap, nil
// }

// func printCities(cities []pokeRes) {
// 	for _, v := range cities {
// 		fmt.Println(v.Name)
// 	}
// }
