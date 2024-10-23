package main

import (
	"testing"
)

// Mock of the CountWords function, replace it with the actual function.
func CountWords(text string) int {
	// mock counting logic
	return len(text) // mock logic, just returns the length of the string
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Empty string", "", 0},
		{"Single word", "word", 4},
		{"Multiple words", "this is a test", 14},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountWords(tc.input)
			if got != tc.expected {
				t.Errorf("CountWords(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
