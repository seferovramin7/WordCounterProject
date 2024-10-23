package utils

import (
	"regexp"
	"strings"
	"unicode"
)

var wordRegex = regexp.MustCompile(`^[a-zA-Z]{3,}$`)

// ProcessText tokenizes and validates text content, returning only valid words.
// Words are considered valid if they are alphabetic and at least 3 characters long.
func ProcessText(content string) []string {
	words := make([]string, 0, len(content)/5) // Estimating an average word length of 5 characters

	var wordBuilder strings.Builder
	for _, r := range content {
		if unicode.IsLetter(r) {
			wordBuilder.WriteRune(unicode.ToLower(r))
		} else if wordBuilder.Len() > 0 { // End of a word
			word := wordBuilder.String()
			if IsValidWord(word) {
				words = append(words, word)
			}
			wordBuilder.Reset() // Reset for the next word
		}
	}

	// Check for the last word if the string didn’t end with a delimiter
	if wordBuilder.Len() > 0 {
		word := wordBuilder.String()
		if IsValidWord(word) {
			words = append(words, word)
		}
	}

	return words
}

// IsValidWord checks if a word is valid (alphabetic and 3 or more characters).
func IsValidWord(word string) bool {
	return wordRegex.MatchString(word)
}
