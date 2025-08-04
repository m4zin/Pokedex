package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/api"
	"pokedex/internal"
	"strings"
)

func main() {
	internal.InitCommands()
	defer api.SharedCache.Cancel()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			internal.CommandExit(nil)
		}

		inputFields := strings.Fields(input)
		cmdName := inputFields[0]
		args := inputFields[1:]

		val, ok := internal.Commands[cmdName]
		if ok {
			if err := val.Callback(args); err != nil {
				fmt.Println(err)
			}
		} else {
			internal.CommandNotFound(nil)
		}
	}
}
