package main

import (
	"fmt"
	"sync"
	"time"
)

func processFile(filename string, wg *sync.WaitGroup) {
	// 2. Signal that this task is finished when the function exits
	defer wg.Done()

	fmt.Printf("Processing %s...\n", filename)
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Printf("Finished %s!\n", filename)
}

func WaitGroupDemo() {
	var wg sync.WaitGroup
	files := []string{"image1.jpg", "image2.png", "image3.gif"}

	for _, file := range files {
		// 1. Increment the counter before starting the goroutine
		wg.Add(1)
		go processFile(file, &wg)
	}

	fmt.Println("Waiting for all files to be processed...")
	
	// 3. Block here until the counter is 0
	wg.Wait()

	fmt.Println("All files processed successfully!")
}