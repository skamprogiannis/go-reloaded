package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			"Hello world",
			[]string{"Hello", "world"},
		},
		{
			"Hello, world!",
			[]string{"Hello", ",", "world", "!"},
		},
		{
			"Wait... what!?",
			[]string{"Wait", "...", "what", "!?"},
		},
		{
			"1E (hex) files",
			[]string{"1E", "(hex)", "files"},
		},
		{
			"word(up)",
			[]string{"word", "(up)"},
		},
		{
			"Line\nBreak",
			[]string{"Line", "\n", "Break"},
		},
		{
			"'Hello'",
			[]string{"'", "Hello", "'"},
		},
	}

	for _, tt := range tests {
		got := tokenize(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("tokenize(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestProcess_Numeric(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1E (hex) files", "30 files"},
		{"10 (bin) years", "2 years"},
		{"Hello world", "Hello world"},
	}

	for _, tt := range tests {
		got := Process(tt.input)
		if got != tt.expected {
			t.Errorf("Process(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestProcess_Style(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello (up)", "HELLO"},
		{"HELLO (low)", "hello"},
		{"hello (cap)", "Hello"},
		{"hELLO (cap)", "Hello"},
		{"one two three (up, 2)", "one TWO THREE"},
		{"one two three (up, 10)", "ONE TWO THREE"}, // Bounds check
	}

	for _, tt := range tests {
		got := Process(tt.input)
		if got != tt.expected {
			t.Errorf("Process(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestProcess_Quotes_Grammar(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"I said ' hello ' .", "I said 'hello'."},
		{"a apple", "an apple"},
		{"A apple", "An apple"},
		{"a Honest mistake", "an Honest mistake"},
		{"a book", "a book"},
		// {"a 'quote'", "a 'quote'"}, // Assumption: quote is not vowel.
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
	}

	for _, tt := range tests {
		got := Process(tt.input)
		if got != tt.expected {
			t.Errorf("Process(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}
