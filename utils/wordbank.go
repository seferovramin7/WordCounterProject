package utils

import (
	"bufio"
	_ "bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

// WordBank is a map that stores valid words.
var WordBank map[string]bool

// LoadWordBank loads words from a local file into the WordBank map.
// The file is expected to contain one word per line, with no empty lines.
func LoadWordBank(filename string) error {
	WordBank = make(map[string]bool)

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		if word != "" && !WordBank[word] {
			WordBank[word] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// DownloadWordsFile downloads the word bank from the URL and stores it locally.
// It uses a larger buffer for more efficient network and disk operations.
func DownloadWordsFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	buf := make([]byte, 32*1024) // 32 KB buffer for faster I/O
	_, err = io.CopyBuffer(out, resp.Body, buf)
	return err
}
