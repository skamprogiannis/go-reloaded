package main

import (
	"testing"
)

func TestAudit(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Case 1: General Casing",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:     "Case 2: Hex and Bin Conversion",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:     "Case 3: Grammar (a -> an)",
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:     "Case 4: Punctuation Groups",
			input:    "Punctuation tests are ... kinda boring ,what do you think ?",
			expected: "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			name:     "Extra 1: Quotes with Punctuation",
			input:    "He said: ' Hello , world (cap) ' .",
			expected: "He said: 'Hello, World'.",
		},
		{
			name:     "Extra 2: Order of Operations",
			input:    "test (up) (low)",
			expected: "test",
		},
		{
			name:     "Extra 3: Case Insensitivity for 'an'",
			input:    "It was a Honest mistake.",
			expected: "It was an Honest mistake.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Process(tt.input)
			if got != tt.expected {
				t.Errorf("\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, got)
			}
		})
	}
}
