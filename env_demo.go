package main

import (
	"fmt"
	"os"
)

func RunEnvDemo() {
	// os.Getenv returns an empty string if the variable is not set
	name := os.Getenv("USER_NAME")
	role := os.Getenv("USER_ROLE")

	if name == "" { name = "Unknown User" }
	if role == "" { role = "Learner" }

	fmt.Printf("--- Docker Env Demo ---\n")
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Your current role is: %s\n", role)

	pw := os.Getenv("DB_PASSWORD")
    
    if pw == "" {
        fmt.Println("Error: DB_PASSWORD not set!")
        return
    }
    
    fmt.Println("Successfully loaded secret of length:", len(pw))
}
