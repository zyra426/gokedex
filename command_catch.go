package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, opts ...string) error {
	if len(opts) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	caught := false

	fmt.Printf("Throwing a Pokeball at %s...\n", opts[0])
	if caught {
		fmt.Printf("%s was caught!\n", opts[0])
		return nil
	}

	fmt.Printf("%s escaped!\n", opts[0])
	return nil
}
