package commands

import (
	"fmt"
	"pokedex/internal"
)

func pokedexCommand(args []string) error {
	if len(internal.Pokedex) > 0 {
		fmt.Println("Your pokedex:")
		for _, v := range internal.Pokedex {
			fmt.Println("-", v.Name)
		}
	}
	return nil
}
