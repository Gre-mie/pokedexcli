package main

import (
	"fmt"
	"strings"
)

// INFO: command structure with methods
type Command struct {
	command string
	arguments []string
	value string
}

func NewCommand(input string) Command {
	tokens := cleanInput(input)
	return Command {
		command: tokens[0],
		// TODO: add aruguments/flags and value
	}
}

func (c Command) Exicute() {
	fmt.Println("Your command was:", c.command) // TEST: prints the command string
}

// INFO: split user input into words
func cleanInput(input string) []string {
	if input == "" {return []string{}}
	clean := strings.Fields(input)
	// lower case all elements in the array, does not strip numbers, punctuation, or other charactors
	newarr := formatArr(clean, strings.ToLower)
	return newarr
}

