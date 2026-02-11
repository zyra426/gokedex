package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, opts ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationsResp.Next
	cfg.prevURL = locationsResp.Previous

	for _, la := range locationsResp.Results {
		fmt.Println(la.Name)
	}

	return nil
}

func commandMapBack(cfg *config, opts ...string) error {
	if cfg.prevURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationsResp.Next
	cfg.prevURL = locationsResp.Previous

	for _, la := range locationsResp.Results {
		fmt.Println(la.Name)
	}
	return nil
}
