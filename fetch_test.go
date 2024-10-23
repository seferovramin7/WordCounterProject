package main

import (
	"errors"
	"testing"
)

// Mock FetchData function for illustration purposes, adapt this to your implementation.
func FetchData(url string) (string, error) {
	if url == "http://valid.url" {
		return "data", nil
	} else {
		return "", errors.New("invalid url")
	}
}

func TestFetchData(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
		wantErr  bool
	}{
		{"Valid URL", "http://valid.url", "data", false},
		{"Invalid URL", "http://invalid.url", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FetchData(tc.url)
			if (err != nil) != tc.wantErr {
				t.Errorf("FetchData() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.expected {
				t.Errorf("FetchData() = %v, want %v", got, tc.expected)
			}
		})
	}
}
