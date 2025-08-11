package main

import (
	buf "bufio"
	f "fmt"
	"os"
)

func main() {
	conf := &config{}

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
			commands[comm].command(conf)
		case "help":
			err := commands[comm].command(conf)
			if err != nil {
				f.Errorf("%v", err)
			}
		case "map":
			err := commands[comm].command(conf)
			if err != nil {
				f.Errorf("%v", err)
			}
		case "mapb":
			err := commands[comm].command(conf)
			if err != nil {
				f.Errorf("%v", err)
			}

		default:
			f.Print("Unknown command\n")
		}
	}

}
