package utils

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// FetchContent retrieves data from a URL and returns it as a single string.
func FetchContent(url string) (string, error) {
	// Optimized Transport for faster connection reuse, keep-alives, and timeout management
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,  // Timeout for connection establishment
			KeepAlive: 30 * time.Second, // Keep-alive to reuse connections
		}).DialContext,
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second, // Time before idle connections are closed
		TLSHandshakeTimeout: 10 * time.Second, // Timeout for TLS handshake
	}

	client := &http.Client{
		Timeout:   10 * time.Second, // Timeout for overall request
		Transport: transport,
	}

	var response *http.Response
	var err error

	// Retry logic with exponential backoff
	for retries := 2; retries > 0; retries-- {
		response, err = client.Get(url)
		if err == nil && response.StatusCode == http.StatusOK {
			break
		}
		if retries == 1 {
			return "", errors.New("failed after multiple retries")
		}
		time.Sleep(time.Duration(4/retries) * time.Second) // Exponential backoff
	}

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read the entire response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil // Return the body as a string
}
