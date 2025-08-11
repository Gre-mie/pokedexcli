package main

import (
	f "fmt"
	"os"
)

type Command struct {
	name string
	description string
	command func(*config) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"exit": Command{
			name: "exit",
			description: "exits the program",
			command: func(conf *config) error {
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
			description: "prints the next 20 locations",
			command: mapForward,
		},
		"mapb": Command{
			name: "mapb",
			description: "prints the last 20 locations",
			command: mapBack,
		},

	}
}

// help - prints the manual for the programme
func printManual(conf *config) error {
	f.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for key, _ := range commands {
		f.Printf("%v: %v\n", key, commands[key].description)
	}
	return nil
}

// map - prints the next 20 locations
func mapForward(conf *config) error {
	
	// testing vvv
	locations, err := getJSON("https://pokeapi.co/api/v2/location-area/") // test
	if err != nil {
		f.Println(err)
		return err
	}
	setConfig(locations, conf)
	// testing ^^^

	return nil
}

// map - prints the last 20 locations
func mapBack(conf *config) error {
	
	// test vvv
	_, err := getJSON("https://pokeapi.co/api/v2/location-area/") // test
	if err != nil {
		f.Println(err)
		return err
	}

	// check if config.last is nil

	// test ^^^


	return nil
}
