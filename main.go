package main

import (
	"time"

	"github.com/zyra426/gokedex/internal/gokeapi"
)

func main() {
	gokeClient := gokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: gokeClient,
	}

	startRepl(cfg)
}
