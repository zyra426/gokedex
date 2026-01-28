package gokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2/location-area/"

type Config struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string
	URL  string
}

func GetMaps(url string) (Config, error) {
	var locationAreas Config
	res, err := http.Get(url)
	if err != nil {
		return locationAreas, fmt.Errorf("error fetching data: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreas, fmt.Errorf("error reading body: %w", err)
	}

	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return locationAreas, fmt.Errorf("error unmarshalling data: %w", err)
	}

	return locationAreas, nil
}
