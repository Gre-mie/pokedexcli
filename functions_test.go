package main

import "testing"

func TestCleanInput(test *testing.T) {
	cases := []struct {
		result string
		expected []string
	}{
		{
			result: "",
			expected: []string{""},
		},
		{
			result: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			result: "Steve Bob James",
			expected: []string{"steve", "bob", "james"},
		},
		{
			result: "@123 4 ABC5 -678",
			expected: []string{"@123", "4", "abc5", "-678"},
		},
	}

	for _, CASE := range cases {
		res := cleanInput(CASE.result)
		if len(res) != len(CASE.expected) {
			test.Errorf("\nexpected: %v\nrecieved: %v", CASE.expected, res)
			continue
		}

		// compare []res recieved array with []expected array
		for i := range res {
			r := res[i]
			e := CASE.expected[i]
			if r == e {
				continue
			} else {
				test.Errorf("\nexpected: %v\nrecieved: %v", CASE.result, res)
				break
			}
		}
		

	}



}
