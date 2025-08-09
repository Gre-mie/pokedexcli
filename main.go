package main

import (
	buf "bufio"
	f "fmt"
	"os"
)

func main() {
	scanner := buf.NewScanner(os.Stdin)
	running := true

	commands := getCommands()

	for running {
		f.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		switch input[0] {
		case "exit":
			commands["exit"].command()
		case "help":
			err := commands["help"].command()
			if err != nil {
				f.Errorf("%v", err)
			}
		default:
			f.Print("Unknown command\n")
		}
	}

}
