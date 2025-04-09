package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon given")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return fmt.Errorf("commandCatch error: %w", err)
	}

	threshold := getThreshold(float64(pokemon.BaseExperience))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	attempt := r.Float64()

	fmt.Println("")
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	// fmt.Println("Threshold: ", threshold)
	// fmt.Println("Attempt: ", attempt)
	if attempt < threshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}

func getThreshold(baseXP float64) float64 {
	minchance := 0.02
	maxchance := 0.80

	maxXP := 700.0
	shaper := 200.0

	threshold := (0.8 - (baseXP/maxXP)) / (1.0 + (baseXP/shaper))

	if threshold > maxchance {
		threshold = maxchance
	}
	if threshold < minchance {
		threshold = minchance
	}
	return threshold
}