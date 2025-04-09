package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/IgorP25/pokedexcli/internal/pokeapi"
)

type config struct {
	pokedex			 map[string]pokeapi.Pokemon
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	commands := registerCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(scanner.Text())
			if len(input) == 0 {
				continue
			}
			args := []string{}
			if len(input) > 1 {
				args = input[1:]
			}

			if command, ok := commands[input[0]]; ok {
				err := command.callback(cfg, args...)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
}

func registerCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			 name:	"help",
			 description: "Displays a help message",
			 callback: commandHelp,
		 },
		 "catch": {
			 name:   "catch <pokemon_name>",
			 description: "Attempt to catch pokemon",
		         callback: commandCatch,
		 },
		 "explore": {
			 name:   "explore <location_name>",
			 description: "Explore a location",
		         callback: commandExplore,
		 },
		 "inspect": {
			 name:   "inspect <pokemon_name>",
			 description: "Inspect previously caught pokemon",
		         callback: commandInspect,
		 },
		"map": {
			 name:   "map",
			 description: "Get the next page of locations",
		         callback: commandMap,
		 },
		 "mapb": {
			 name:   "mapb",
			 description: "Get the previous page of locations",
		         callback: commandMapb,
		 },
		 "pokedex": {
			 name:   "pokedex",
			 description: "Display list of all caught Pokemon",
		         callback: commandPokedex,
		 },
		 "exit": {
			 name:   "exit",
			 description: "Exit the Pokedex",
		         callback: commandExit,
		 },
	 }
 }
 