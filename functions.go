package main

import str "strings"

// lower cases input and splits it into an array of strings
func cleanInput(input string) []string {
	words := str.Split(str.ToLower(input), " ")
	return words
}