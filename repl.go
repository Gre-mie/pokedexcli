package main

import (
	"strings"
)

// split user input into words
func cleanInput(input string) []string {
	if input == "" {return []string{}}
	clean := strings.Fields(input)
	// lower case all elements in the array
	newarr := formatArr(clean, strings.ToLower)
	return newarr
}

