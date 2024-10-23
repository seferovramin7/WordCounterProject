package main

import (
	"testing"
)

// Mock WordBank struct for illustration.
type WordBank struct {
	words []string
}

func (wb *WordBank) AddWord(word string) {
	wb.words = append(wb.words, word)
}

func (wb *WordBank) GetWords() []string {
	return wb.words
}

func TestWordBank(t *testing.T) {
	wb := &WordBank{}

	t.Run("AddWord", func(t *testing.T) {
		wb.AddWord("hello")
		if len(wb.GetWords()) != 1 || wb.GetWords()[0] != "hello" {
			t.Errorf("Expected 'hello' in word bank, got %v", wb.GetWords())
		}
	})

	t.Run("GetWords", func(t *testing.T) {
		words := wb.GetWords()
		if len(words) != 1 || words[0] != "hello" {
			t.Errorf("Expected 'hello', got %v", words)
		}
	})
}
