package main

import (
	"bufio"
	"fmt"
	"os"

	"internal/pokedexAPI"
)

type command struct {
	name        string
	description string
	callback    func(config *pokedexAPI.Config) error
}

func commandHelp(_ *pokedexAPI.Config) error {
	commands := generateCommandMap()
	fmt.Print("Command usage:\n\n")
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(_ *pokedexAPI.Config) error {
	os.Exit(0)
	return nil
}

func commandMap(config *pokedexAPI.Config) error {
	if config.Next == "" {
		*config = pokedexAPI.GetAreas(nil)
	} else {
		*config = pokedexAPI.GetAreas(&config.Next)
	}
	return nil
}

func commandMapb(config *pokedexAPI.Config) error {
	if config.Previous == "" {
		*config = pokedexAPI.GetAreas(nil)
	} else {
		*config = pokedexAPI.GetAreas(&config.Previous)
	}
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
		"map": {
			name:        "map",
			description: " diplays the next set of nearby locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous set of nearby locations",
			callback:    commandMapb,
		},
	}
}

func main() {
	commandMap := generateCommandMap()
	var config *pokedexAPI.Config = &pokedexAPI.Config{
		Next:     "",
		Previous: "",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command, ok := commandMap[input]
		if ok {
			command.callback(config)
		} else {
			fmt.Println("command not found")
			break
		}
	}
}
