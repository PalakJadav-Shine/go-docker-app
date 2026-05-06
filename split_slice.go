package main

import (
	"fmt"
	"strings"
)

func SliceDemo() {
	fmt.Println("--- Split vs Slice Demo ---")

	// 1. SPLIT (The Action)
	// strings.Split takes a string and turns it into a Slice of strings
	rawString := "Go,Docker,CLI,Colima"
	splitResult := strings.Split(rawString, ",")
	fmt.Printf("Split result (type %%T): %T\n", splitResult)
	fmt.Println("Items:", splitResult)

	// 2. SLICE (The Structure)
	// A slice is a view into an underlying array.
	// Let's take a "slice" of our split result
	mySubSlice := splitResult[1:3] // Grabs index 1 and 2 (Docker and CLI)
	
	fmt.Println("\nFull Slice:", splitResult)
	fmt.Println("Sub-slice [1:3]:", mySubSlice)
}
