package main

import "fmt"

func main1() {
    fmt.Println("=== STARTING ALL SERVICES SEQUENTIALLY ===")

	RunEnvDemo()

	fmt.Println("------------------------------------------")

	RunConcurrencyDemo()

    fmt.Println("------------------------------------------")

    HealthCheckDemo()

    fmt.Println("------------------------------------------")

    SliceDemo()

    fmt.Println("------------------------------------------")

    WaitGroupDemo()

    fmt.Println("------------------------------------------")

    HttpDemo()  // This will block the main goroutine, so it should be last in the sequence

	fmt.Println("=== ALL TASKS FINISHED ===")

}
