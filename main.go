package main

import (
	"bufio"
	"fmt"
	"os"
)

// INFO: main function automatically called
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	running := true
	for running {
		// INFO: scan for input
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		userinput := scanner.Text()
		err := scanner.Err()
		if err != nil {
			running = false
			fmt.Println(err)
		}
		
		// INFO: evaluate input
		if userinput == "" {continue}
		command := NewCommand(userinput)

		// INFO: preform operations input -> output

		// INFO: print output 
		command.Exicute() // TEST: 
	}

}

