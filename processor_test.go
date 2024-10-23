package main

import (
	"testing"
)

// A simple mock function for illustration, adapt it to your actual `ProcessData` function.
func ProcessData(input string) string {
	// mock processing logic
	return input + " processed"
}

func TestProcessData(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", " processed"},
		{"Normal input", "data", "data processed"},
		{"Long input", "long input string", "long input string processed"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ProcessData(tc.input)
			if got != tc.expected {
				t.Errorf("ProcessData(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
