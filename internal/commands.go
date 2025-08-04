package internal

import (
	"fmt"
	"log"
	"os"
	"pokedex/api"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(args []string) error
}

var Commands = make(map[string]cliCommand)
var currLocationAreasUrl = api.PokeApiBaseUrl + api.LocationAreaPath
var previous string
var currLocations []string

func commandHelp(args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range Commands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func commandMap(args []string) error {
	data, err := api.GetLocationAreas(currLocationAreasUrl)
	if err != nil {
		log.Fatalf("error fetching location areas: %v", err)
	}

	var tempLocations []string
	for _, v := range data.Results {
		tempLocations = append(tempLocations, v.Name)
		fmt.Println(v.Name)
	}
	currLocations = tempLocations
	if data.Previous != nil {
		previous = *data.Previous
	} else {
		previous = api.PokeApiBaseUrl + api.LocationAreaPath
	}
	currLocationAreasUrl = data.Next
	return nil
}

func commandMapB(args []string) error {
	if previous == "" {
		fmt.Println("no previous locations to be found!")
		return nil
	}
	data, err := api.GetLocationAreas(previous)
	if err != nil {
		log.Fatalf("error fetching location areas: %v", err)
	}
	var tempLocations []string
	for _, v := range data.Results {
		tempLocations = append(tempLocations, v.Name)
		fmt.Println(v.Name)
	}
	currLocations = tempLocations
	return nil
}

func commandExplore(args []string) error {
	if len(args) != 1 {
		fmt.Println("Expecting a single location area")
		return nil
	}
	if len(currLocations) == 0 {
		fmt.Println("Run map first to get list of locations")
		return nil
	}
	for _, v := range currLocations {
		if v == args[0] {
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
	}
	return fmt.Errorf("could not find location area in map")
}

func CommandNotFound(args []string) {
	fmt.Println("Could not find command, try again")
}

func CommandExit(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func InitCommands() {
	Commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		Callback:    CommandExit,
	}
	Commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		Callback:    commandHelp,
	}
	Commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		Callback:    commandMap,
	}
	Commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of previous 20 location areas in the Pokemon world",
		Callback:    commandMapB,
	}
	Commands["explore"] = cliCommand{
		name:        "explore",
		description: "Displays the pokemons present in a location area",
		Callback:    commandExplore,
	}
}
