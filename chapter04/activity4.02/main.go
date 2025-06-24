package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userID := os.Args[1]

	names := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}

	name, exists := names[userID]

	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", userID)

		for key, value := range names {
			fmt.Println(" ID:", key, "Name:", value)
		}
		os.Exit(1)
	}

	fmt.Println("Name:", name)
}
