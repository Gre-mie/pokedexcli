package main

import (
	"testing"
	"reflect"
//	"strings"
	"fmt"
)

func format [Te any, Ta any] (name string, expected Te, actual Ta) string { 
	return fmt.Sprintf(`
test: %s
	expected: %v
	actual:   %v`, name, expected, actual)
}

// INFO: cleanInput should return an array of words from a string
func TestCleanInput(t *testing.T) {
	testCases := []struct{
		name string
		input string
		output []string
	}{
		{
			name: "Empty input",
			input: "",
			output: []string{},
		},
		{
			name: "Punctuation",
			input: "Hello World!",
			output: []string{"hello", "world!"},
		},
		{
			name: "Multiple whitespace charactors",
			input: "White   Space\tChars\n",
			output: []string{"white", "space", "chars"},
		},
		{
			name: "Multi line input",
			input: "multi\nline\ninput",
			output: []string{"multi", "line", "input"},
		},
	}

	for _, test := range testCases {
		var compare []string = cleanInput(test.input)
			reflect.DeepEqual(compare, test.output)
	}

}

