package main

import (
	"errors"
	"fmt"
)

func wrapMap(cfg *config, isPrev bool, args ...string) error {
	var url *string
	if isPrev {
		if cfg.prevLocationsURL == nil {
			return errors.New("you're on the first page")
		} else {
			url = cfg.prevLocationsURL
		}
	} else {
		url = cfg.nextLocationsURL
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return fmt.Errorf("wrapMap error: %w", err)
	}
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}
	return nil
}
func commandMap(cfg *config, args ...string) error {
	err := wrapMap(cfg, false, args...)
	if err != nil {
		return fmt.Errorf("commandMap error: %w", err)
	}
	return nil
}
func commandMapb(cfg *config, args ...string) error {
	err := wrapMap(cfg, true, args...)
	if err != nil {
		return fmt.Errorf("commandMapb error: %w", err)
	}
	return nil
}
