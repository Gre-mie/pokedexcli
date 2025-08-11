package main

import (
	f "fmt"
	"os"
)

type Command struct {
	name string
	description string
	command func() error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"exit": Command{
			name: "exit",
			description: "exits the program",
			command: func() error {
				f.Println("Closing the Pokedex... Goodbye!")
				os.Exit(0)
				return nil
			},
		},
		"help": Command{
			name: "help",
			description: "prints a help message describing how to use the REPL",
			command: printManual,
		},
		"map": Command{
			name: "map",
			description: "",
			command: func() error {
				f.Println("map command is running")
				return nil
			},
		},

	}
}

func printManual() error {
	f.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for key, _ := range commands {
		f.Printf("%v: %v\n", key, commands[key].description)
	}
	return nil
}
