package commands

import (
	"fmt"
	"os"
)

func helpCommand(args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range Commands {
		fmt.Printf("%v: %v\n", v.Name, v.Description)
	}
	return nil
}

func ExitCommand(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func NotFoundCommand(args []string) {
	fmt.Println("Could not find command, try again")
}
