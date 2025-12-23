package main

import (
	"testing"
	"strings"
		"fmt"
)

func format [Te any, Ta any] (name string, expected Te, actual Ta) string { 
	return fmt.Sprintf(`
test: %s
	expected: %v
	actual:   %v`, name, expected, actual)
}

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
		if len(compare) != len(test.output) {
			t.Errorf(format(test.name, len(test.output), len(compare)))
			continue
		}

		// check arrays match
		for i:=0; i<len(test.output); i++ {
			expected := test.output[i]
			word := compare[i]
			if strings.Compare(expected, word) != 0 {
				t.Errorf(format(test.name, test.output, compare))
				break
			}	
		}
	}

}

