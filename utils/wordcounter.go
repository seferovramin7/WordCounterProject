package utils

import (
	"sort"
	"sync"
	"sync/atomic"
)

// WordCounter is a structure that keeps track of word counts with thread-safety.
type WordCounter struct {
	counts sync.Map // Using sync.Map for concurrent access
}

// NewWordCounter creates a new WordCounter instance.
func NewWordCounter() *WordCounter {
	return &WordCounter{}
}

// AddWords increments the count of each word found.
func (wc *WordCounter) AddWords(words []string) {
	localCounts := make(map[string]int)

	// Count words locally first to reduce lock contention
	for _, word := range words {
		localCounts[word]++
	}

	// Update the shared word counter with the local counts
	for word, count := range localCounts {
		wc.incrementWordCount(word, count)
	}
}

// incrementWordCount safely increments the word count in the sync.Map
func (wc *WordCounter) incrementWordCount(word string, count int) {
	// Try to load the value from the map
	actual, loaded := wc.counts.LoadOrStore(word, new(int64)) // Store a pointer to int64

	if loaded {
		// If loaded, increment the existing value using atomic
		atomic.AddInt64(actual.(*int64), int64(count))
	} else {
		// Otherwise, initialize the count
		atomic.StoreInt64(actual.(*int64), int64(count))
	}
}

// GetTopWords returns the top N most frequent words.
func (wc *WordCounter) GetTopWords(n int) []map[string]int {
	wordList := wc.getWordList()

	// Sort the slice by frequency (descending)
	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i].Count > wordList[j].Count
	})

	// Select top N words
	var topWords []map[string]int
	for i := 0; i < n && i < len(wordList); i++ {
		topWords = append(topWords, map[string]int{wordList[i].Word: wordList[i].Count})
	}

	return topWords
}

// wordFreq holds a word and its frequency count.
type wordFreq struct {
	Word  string
	Count int
}

// getWordList converts the sync.Map to a slice of wordFreq for sorting.
func (wc *WordCounter) getWordList() []wordFreq {
	var wordList []wordFreq

	// Convert the sync.Map to a slice of word-frequency pairs
	wc.counts.Range(func(key, value interface{}) bool {
		wordList = append(wordList, wordFreq{
			Word:  key.(string),
			Count: int(atomic.LoadInt64(value.(*int64))), // Load the int64 value atomically
		})
		return true
	})

	return wordList
}
