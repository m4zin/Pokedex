package commands

import (
	"fmt"
	"pokedex/api"
)

func exploreCommand(args []string) error {
	if len(args) != 1 {
		fmt.Println("Expecting a single location area")
		return nil
	}
	fmt.Println("Exploring " + args[0] + "...")
	data, err := api.GetPokemonsAtLocation(api.PokeApiBaseUrl + api.LocationAreaPath + "/" + args[0])
	if err != nil {
		return fmt.Errorf("error fetching pokemons at location: %w", err)
	}
	fmt.Println("Found Pokemon:")
	for _, v := range data.PokemonEncounters {
		fmt.Println("-", v.Pokemon.Name)
	}
	return nil
}
