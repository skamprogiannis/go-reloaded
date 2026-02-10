package main

import (
	"os"
	"testing"
)

func TestCLI_BasicCopy(t *testing.T) {
	// Setup
	inputContent := "Hello World"
	inputFile := "test_input.txt"
	outputFile := "test_output.txt"

	// Ensure clean state
	os.Remove(inputFile)
	os.Remove(outputFile)

	err := os.WriteFile(inputFile, []byte(inputContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test input file: %v", err)
	}
	defer os.Remove(inputFile)
	defer os.Remove(outputFile)

	// Mock args
	os.Args = []string{"cmd", inputFile, outputFile}

	// Execute
	main()

	// Verify
	outputContent, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read test output file: %v", err)
	}

	if string(outputContent) != inputContent {
		t.Errorf("Expected '%s', got '%s'", inputContent, string(outputContent))
	}
}
