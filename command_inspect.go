package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you haven't caught this pokemon")
	}

	fmt.Printf("Name: %s", pokemon.Name)
	fmt.Printf("Height: %v", pokemon.Height)
	fmt.Printf("Weight: %v", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v", stat.Stat.Name, stat.BaseStat)
	}
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s", typ.Type.Name)
	}

	return nil
}
