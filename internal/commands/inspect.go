package commands

import (
	"fmt"
	"pokedex/helper"
	"pokedex/internal"
)

func inspectCommand(args []string) error {
	if len(args) != 1 {
		fmt.Println("Expecting a pokemon")
		return nil
	}
	val, ok := internal.Pokedex[args[0]]
	if ok {
		helper.PrettyPrint(val)
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
