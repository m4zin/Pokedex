package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func(args []string) error
}

var Commands = make(map[string]CliCommand)

func InitCommands() {
	Commands = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    ExitCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    helpCommand,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    MapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of previous 20 location areas in the Pokemon world",
			Callback:    MapBackCommand,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the pokemons present in a location area",
			Callback:    exploreCommand,
		},
		"catch": {
			Name:        "catch",
			Description: "Catches a pokemon",
			Callback:    catchCommand,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspects a caught pokemon",
			Callback:    inspectCommand,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Prints all the caught pokemons",
			Callback:    pokedexCommand,
		},
	}
}
