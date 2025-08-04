package internal

import (
	"fmt"
	"log"
	"os"
	"pokedex/api"
	"pokedex/helper"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(args []string) error
}

var pokedex = make(map[string]api.Pokemon)
var Commands = make(map[string]cliCommand)
var currLocationAreasUrl = api.PokeApiBaseUrl + api.LocationAreaPath
var previous string

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

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
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
	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandExplore(args []string) error {
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

func commandCatch(args []string) error {
	if len(args) != 1 {
		fmt.Println("Expecting a pokemon")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	data, err := api.GetPokemon(api.PokeApiBaseUrl + api.PokemonPath + "/" + args[0])
	if err != nil {
		return fmt.Errorf("error fetching pokemon: %w", err)
	}
	attempt := helper.AttemptCatch(data.BaseExperience)
	if attempt {
		pokedex[args[0]] = data
		fmt.Printf("%v was caught!\n", args[0])
	} else {
		fmt.Printf("%v escaped!\n", args[0])
	}
	return nil
}

func commandInspect(args []string) error {
	if len(args) != 1 {
		fmt.Println("Expecting a pokemon")
		return nil
	}
	val, ok := pokedex[args[0]]
	if ok {
		helper.PrettyPrint(val)
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(args []string) error {
	if len(pokedex) > 0 {
		fmt.Println("Your pokedex:")
		for _, v := range pokedex {
			fmt.Println("-", v.Name)
		}
	}
	return nil
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
	Commands["catch"] = cliCommand{
		name:        "catch",
		description: "Catches a pokemon",
		Callback:    commandCatch,
	}
	Commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspects a caught pokemon",
		Callback:    commandInspect,
	}
	Commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Prints all the caught pokemons",
		Callback:    commandPokedex,
	}
}
