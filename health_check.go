package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkWebsite(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when the function ends

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[ERROR] %s is down!\n", url)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Round(time.Millisecond)
	fmt.Printf("[SUCCESS] %s returned %d in %s\n", url, resp.StatusCode, elapsed)
}

func HealthCheckDemo() {
	websites := []string{
		"https://google.com",
		"https://github.com",
		"https://go.dev",
		"https://docker.com",
	}

	var wg sync.WaitGroup

	fmt.Println("Starting concurrent health checks...")
	totalStart := time.Now()

	for _, url := range websites {
		wg.Add(1)            // Increment counter
		go checkWebsite(url, &wg) // Start Goroutine
	}

	wg.Wait() // Wait for all counter to hit zero

	fmt.Printf("\nDone! Total time for all checks: %s\n", time.Since(totalStart).Round(time.Millisecond))
}