package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)

		if len(actual) != len(cs.expected) {
			t.Errorf("Length mismatch: got %v, expected %v", len(actual), len(cs.expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("Mismatch at index %d: got '%s', expected '%s'", i, actualWord, expectedWord)
			}
		}
	}
}
