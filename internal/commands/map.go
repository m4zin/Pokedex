package commands

import (
	"fmt"
	"log"
	"pokedex/api"
	"pokedex/internal"
)

func MapCommand(args []string) error {
	data, err := api.GetLocationAreas(internal.CurrLocationAreasUrl)
	if err != nil {
		return fmt.Errorf("error fetching location areas: %v", err)
	}
	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	if data.Previous != nil {
		internal.Previous = *data.Previous
	} else {
		internal.Previous = api.PokeApiBaseUrl + api.LocationAreaPath
	}
	internal.CurrLocationAreasUrl = data.Next
	return nil
}

func MapBackCommand(args []string) error {
	if internal.Previous == "" {
		fmt.Println("no previous locations to be found!")
		return nil
	}
	data, err := api.GetLocationAreas(internal.Previous)
	if err != nil {
		log.Fatalf("error fetching location areas: %v", err)
	}
	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	return nil
}
