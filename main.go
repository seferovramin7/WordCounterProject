package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/seferovramin7/WordCounterProject/utils"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

const workerCount = 15 // Number of Goroutines for word processing

// setupProfiling sets up CPU profiling and ensures it stops when the program ends.
func setupProfiling() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile() // Stop profiling when the program ends
}

// loadURLs loads URLs from a file and returns a slice of URLs.
func loadURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", filename, err)
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filename, err)
	}
	return urls, nil
}

func main() {
	// Step 1: Start CPU profiling
	setupProfiling()

	// Step 2: Record the start time
	startTime := time.Now()

	// Step 3: Load URLs from file
	urls, err := loadURLs("endg-urls")
	if err != nil {
		log.Fatal("Error loading URLs:", err)
	}

	// Step 4: Download and load word bank
	wordBankURL := "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt"
	if err := utils.DownloadWordsFile(wordBankURL, "words.txt"); err != nil {
		log.Fatalf("Error downloading word bank from %s: %v", wordBankURL, err)
	}
	if err := utils.LoadWordBank("words.txt"); err != nil {
		log.Fatal("Error loading word bank:", err)
	}

	// Step 5: Create a word counter and a results channel
	wordCounter := utils.NewWordCounter()
	resultChannel := make(chan string)

	// Step 6: Launch a worker pool for concurrent word processing
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for content := range resultChannel {
				words := utils.ProcessText(content) // Tokenize and validate words
				wordCounter.AddWords(words)         // Add words to the shared counter
			}
		}()
	}

	// Step 7: Fetch data from URLs and send content to the resultChannel
	var fetchWG sync.WaitGroup
	for _, url := range urls {
		fetchWG.Add(1)
		go func(u string) {
			defer fetchWG.Done()
			content, err := utils.FetchContent(u) // Fetch content from the URL
			if err != nil {
				log.Printf("Error fetching content from %s: %v", u, err)
				return
			}
			resultChannel <- content // Send fetched content to be processed
		}(url)
	}

	// Close the result channel once all URLs are fetched
	fetchWG.Wait()
	close(resultChannel)

	// Wait for all workers to finish processing
	wg.Wait()

	// Step 8: Get the top 10 most frequent words
	topWords := wordCounter.GetTopWords(10)

	// Step 9: Print the top words in a JSON formatted output
	jsonResult, err := json.MarshalIndent(topWords, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON result: %v", err)
	}
	fmt.Println(string(jsonResult))

	// Step 10: Calculate and print the elapsed time
	elapsedTime := time.Since(startTime)
	fmt.Printf("Process completed in %s\n", elapsedTime)
}
