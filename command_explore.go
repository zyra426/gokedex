package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, opts ...string) error {
	if len(opts) != 1 {
		return errors.New("you must provide a location name")
	}

	exploreResp, err := cfg.pokeapiClient.ListAreaInfo(opts[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", exploreResp.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range exploreResp.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
