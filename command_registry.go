package main

import (
	"fmt"
	"os"
)

type config struct {
	commands map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	message := "Welcome to the Pokedex!\nUsage:\n\n"

	for name, command := range cfg.commands {
		message += name + ": " + command.description + "\n"
	}
	fmt.Print(message)
	return nil
}

func commandMap(cfg *config) error {

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
