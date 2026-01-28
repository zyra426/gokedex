package main

import "testing"

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}

	cases := []testCase{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "caught in the eye of the storm",
			expected: []string{"caught", "in", "the", "eye", "of", "the", "storm"},
		},
		{
			input:    "weapon   of choosing    ",
			expected: []string{"weapon", "of", "choosing"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("got: %v, expected: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]

			if word != expected {
				t.Errorf("word: %s, expected: %s", word, expected)
			}
		}
	}
}
