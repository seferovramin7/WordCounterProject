package main

import (
	"testing"
)

// Example utility function in `main.go`
func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	if got := Sum(1, 2); got != 3 {
		t.Errorf("Sum(1, 2) = %v; want 3", got)
	}
}
