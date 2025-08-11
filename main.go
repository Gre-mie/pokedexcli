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

		comm := input[0]
		switch comm {
		case "exit":
			commands[comm].command()
		case "help":
			err := commands[comm].command()
			if err != nil {
				f.Errorf("%v", err)
			}
		case "map":
			err := commands[comm].command()
			if err != nil {
				f.Errorf("%v", err)
			}
		default:
			f.Print("Unknown command\n")
		}
	}

}
