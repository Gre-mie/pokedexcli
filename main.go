package main

import (
	buf "bufio"
	f "fmt"
	"os"
)

func main() {
	scanner := buf.NewScanner(os.Stdin)
	running := true

	for running {
		f.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		f.Printf("Your command was: %v\n", input[0])
	}

}
