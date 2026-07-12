package main

import "testing"

func cleanTest(t *testing.T) {
	testCase := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello there ",
			expected: []string{"hello", "there"},
		},
		{
			input:    "h j km e",
			expected: []string{"h", "j", "km", "e"},
		},
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
	}

	for i := range testCase {
		actual := cleanInput(testCase[i].input)

		for j, e := range testCase[i].expected {
			if actual[j] != e {
				t.Errorf("Cleaning failed")
			}
		}

	}
}
