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

func main() {
	// Step 1: Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile() // Stop profiling when the program ends

	// Step 2: Record the start time
	startTime := time.Now()

	// Step 2: Load URLs from file
	urls, err := loadURLs("endg-urls")
	if err != nil {
		log.Fatal("Error loading URLs:", err)
	}

	// Step 3: Download and load word bank
	wordBankURL := "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt"
	if err := utils.DownloadWordsFile(wordBankURL, "words.txt"); err != nil {
		log.Fatal("Error downloading word bank:", err)
	}
	if err := utils.LoadWordBank("words.txt"); err != nil {
		log.Fatal("Error loading word bank:", err)
	}

	// Step 4: Create a word counter and results channel
	wordCounter := utils.NewWordCounter()
	resultChannel := make(chan string)

	// Step 5: Launch a worker pool for processing words concurrently
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

	// Step 6: Fetch data from URLs and send the content to the resultChannel
	var fetchWG sync.WaitGroup
	for _, url := range urls {
		fetchWG.Add(1)
		go func(u string) {
			defer fetchWG.Done()
			content, err := utils.FetchContent(u) // Fetch content
			if err != nil {
				log.Println("Error fetching URL:", err)
				return
			}
			resultChannel <- content // Send fetched content for processing
		}(url)
	}

	// Close the result channel once all URLs are fetched
	fetchWG.Wait()
	close(resultChannel)

	// Wait for all workers to finish processing
	wg.Wait()

	// Step 7: Get top 10 words
	topWords := wordCounter.GetTopWords(10)

	// Step 8: Print JSON result
	jsonResult, _ := json.MarshalIndent(topWords, "", "  ")
	fmt.Println(string(jsonResult))

	// Step 9: Calculate and print elapsed time
	elapsedTime := time.Since(startTime)
	fmt.Printf("Process completed in %s\n", elapsedTime)
}

func loadURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return urls, nil
}
