package main

import (
	"github.com/seferovramin7/WordCounterProject/utils"
	"testing"
)

// Benchmark the ProcessText function
func BenchmarkProcessText(b *testing.B) {
	sampleContent := "This is a simple test content for benchmarking. We will run this through the process function."

	// Load the word bank before the benchmark if necessary
	wordBankURL := "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt"
	utils.DownloadWordsFile(wordBankURL, "words.txt")
	utils.LoadWordBank("words.txt")

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		utils.ProcessText(sampleContent)
	}
}
