package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zyra426/gokedex/internal/gokeapi"
)

type config struct {
	pokeapiClient gokeapi.Client
	nextURL       *string
	prevURL       *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}

		command := text[0]
		options := []string{}
		if len(text) > 1 {
			options = text[1:]
		}

		cliComm, exists := supportedCommands()[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := cliComm.callback(cfg, options...)
		if err != nil {
			fmt.Println(err)
		}
		continue
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func supportedCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays list of maps",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous list of maps",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Displays list of pokemon in the requested area, e.g. explore <area-name>",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
