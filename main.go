package main

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	commands := generateCommandMap()
	fmt.Print("Command usage:\n\n")
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func generateCommandMap() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	commandMap := generateCommandMap()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command, ok := commandMap[input]
		if ok {
			command.callback()
		} else {
			fmt.Println("command not found")
			break
		}
	}
}
