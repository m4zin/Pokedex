package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/api"
	"pokedex/internal/commands"
	"strings"
)

func repl() {
	defer api.SharedCache.Cancel()
	commands.InitCommands()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			commands.ExitCommand(nil)
		}

		inputFields := strings.Fields(input)
		cmdName := inputFields[0]
		args := inputFields[1:]

		val, ok := commands.Commands[cmdName]
		if ok {
			if err := val.Callback(args); err != nil {
				fmt.Println(err)
			}
		} else {
			commands.NotFoundCommand(nil)
		}
	}
}
