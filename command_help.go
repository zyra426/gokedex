package main

import (
	"fmt"

	"github.com/zyra426/gokedex/internal/gokeapi"
)

func commandHelp(cfg *gokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range supportedCommands() {
		fmt.Printf("%s: %s", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
