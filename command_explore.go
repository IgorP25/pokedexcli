package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location given")
	}

	location, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return fmt.Errorf("commandExplore error: %w", err)
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range location.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}
	return nil
}