package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"internal/pokedexAPI"
)

type command struct {
	name        string
	description string
	callback    func(config *pokedexAPI.Config, area string) error
}

func commandHelp(_ *pokedexAPI.Config, _ string) error {
	commands := generateCommandMap()
	fmt.Print("Command usage:\n\n")
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(_ *pokedexAPI.Config, _ string) error {
	os.Exit(0)
	return nil
}

func commandMap(config *pokedexAPI.Config, _ string) error {
	if config.Next == "" {
		*config = pokedexAPI.GetAreas(nil)
	} else {
		*config = pokedexAPI.GetAreas(&config.Next)
	}
	return nil
}

func commandMapb(config *pokedexAPI.Config, _ string) error {
	if config.Previous == "" {
		*config = pokedexAPI.GetAreas(nil)
	} else {
		*config = pokedexAPI.GetAreas(&config.Previous)
	}
	return nil
}

func commandExplore(_ *pokedexAPI.Config, area string) error {
	pokedexAPI.GetAreaPokemon(area)
	return nil
}

func commandCatch(_ *pokedexAPI.Config, name string) error {
	pokedexAPI.CatchPokemon(name)
	return nil
}

func commandInspect(_ *pokedexAPI.Config, name string) error {
	pokedexAPI.InspectPokemon(name)
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
		"explore": {
			name:        "explore",
			description: "displays Pokemon that can be found in area provided",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempts to catch the provided pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "displays the name, height, weight, stats, and type of provided pokemon",
			callback:    commandInspect,
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

		commandList := strings.Fields(input)
		command, ok := commandMap[commandList[0]]
		if ok {
			args := command.name == "explore"
			args = args || command.name == "catch"
			args = args || command.name == "inspect"
			if args {
				command.callback(config, commandList[1])
			} else {
				command.callback(config, "")
			}
		} else {
			fmt.Println("command not found")
			break
		}
	}
}
