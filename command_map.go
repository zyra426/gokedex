package main

import (
	"fmt"

	"github.com/zyra426/gokedex/internal/gokeapi"
)

func commandMap(cfg *gokeapi.Config) error {
	if cfg.Next != "" {
		locationAreas, err := gokeapi.GetMaps(cfg.Next)
		if err != nil {
			return err
		}

		for _, la := range locationAreas.Results {
			fmt.Println(la.Name)
		}

		return nil
	}

	locationAreas, err := gokeapi.GetMaps(gokeapi.BaseURL)
	if err != nil {
		return err
	}

	for _, la := range locationAreas.Results {
		fmt.Println(la.Name)
	}

	return nil
}
