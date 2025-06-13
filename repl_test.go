package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello world Y  ",
			expected: []string{"hello", "world", "y"},
		},
		{
			input:    "",
			expected: []string{""},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)

		if len(result) != len(c.expected) {
			t.Errorf("lenght didn't match for string: %s", c.input)
			return
		}

		for i := range result {
			word := result[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word %s didn't match expected word %s in test string %s", word, expectedWord, c.input)
				return
			}

		}

	}
}
