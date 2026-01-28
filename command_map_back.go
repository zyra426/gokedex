package main

import (
	"fmt"

	"github.com/zyra426/gokedex/internal/gokeapi"
)

func commandMapBack(cfg *gokeapi.Config) error {
	if cfg.Previous == "" {
		fmt.Println("You're on the first page")
	}

	locationAreas, err := gokeapi.GetMaps(cfg.Previous)
	if err != nil {
		fmt.Println("Cannot retrieve previous page")
	}

	for _, la := range locationAreas.Results {
		fmt.Println(la.Name)
	}
	return nil
}
