package commands

import (
	"fmt"
	"pokedex/api"
	"pokedex/helper"
	"pokedex/internal"
)

func catchCommand(args []string) error {
	if len(args) != 1 {
		fmt.Println("Invalid catch usage")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	data, err := api.GetPokemon(api.PokeApiBaseUrl + api.PokemonPath + "/" + args[0])
	if err != nil {
		return fmt.Errorf("error fetching pokemon: %w", err)
	}
	attempt := helper.AttemptCatch(data.BaseExperience)
	if attempt {
		internal.Pokedex[args[0]] = data
		fmt.Printf("%v was caught!\n", args[0])
	} else {
		fmt.Printf("%v escaped!\n", args[0])
	}
	return nil
}
