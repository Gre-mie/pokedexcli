package main

import str "strings"

func cleanInput(input string) []string {
	words := str.Split(str.ToLower(input), " ")
	return words
}