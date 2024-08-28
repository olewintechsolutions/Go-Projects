package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the user provided a name as a command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}

	// Get the name from the command line arguments
	name := os.Args[1]

	// Print a personalized greeting
	fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)
}
