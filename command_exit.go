package main

import (
	"fmt"
	"os"

	"github.com/zyra426/gokedex/internal/gokeapi"
)

func commandExit(cfg *gokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
