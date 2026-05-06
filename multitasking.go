package main

import (
	"fmt"
	"time"
)

func performTask(id int, ch chan string) {
	// Simulate a task taking time
	fmt.Printf("Task %d started...\n", id)
	time.Sleep(2 * time.Second) 
	
	// Send the result back through the channel
	ch <- fmt.Sprintf("Task %d is complete!", id)
}

func RunConcurrencyDemo() {
	// 1. Create a channel to communicate
	resultsChannel := make(chan string)

	// 2. Start 3 tasks concurrently using 'go' keyword
	for i := 1; i <= 3; i++ {
		go performTask(i, resultsChannel)
	}

	fmt.Println("Waiting for tasks to finish...")

	// 3. Collect the results from the channel
	// This part "blocks" (waits) until data is sent into the channel
	for i := 1; i <= 3; i++ {
		result := <-resultsChannel
		fmt.Println(result)
	}

	fmt.Println("All tasks finished. Total time elapsed: ~2 seconds (not 6!)")
}