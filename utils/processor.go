package utils

import (
	"regexp"
	"strings"
	"unicode"
)

var wordRegex = regexp.MustCompile(`^[a-zA-Z]{3,}$`)

// ProcessText tokenizes and validates text content.
func ProcessText(content string) []string {
	// Pre-allocate memory based on a rough estimate of the number of words
	words := make([]string, 0, len(content)/5) // Estimating an average word length of 5 characters

	// Convert content to lowercase and iterate through characters
	var wordBuilder strings.Builder
	for _, r := range content {
		if unicode.IsLetter(r) {
			wordBuilder.WriteRune(unicode.ToLower(r))
		} else if wordBuilder.Len() > 0 { // End of a word
			word := wordBuilder.String()
			if len(word) >= 3 && IsValidWord(word) { // Simple check on word length
				words = append(words, word)
			}
			wordBuilder.Reset() // Reset for the next word
		}
	}

	// Check for the last word if the string didn't end with a delimiter
	if wordBuilder.Len() > 0 {
		word := wordBuilder.String()
		if len(word) >= 3 && IsValidWord(word) {
			words = append(words, word)
		}
	}

	return words
}
