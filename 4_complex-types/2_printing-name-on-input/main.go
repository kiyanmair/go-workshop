package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Must provide exactly one argument, the user ID.")
		os.Exit(1)
	}
	id := os.Args[1]
	name, exists := users[id]
	if !exists {
		fmt.Printf("Error: User with ID = %v does not exist.\n", id)
		os.Exit(1)
	}
	fmt.Printf("Hi, %v\n", name)
}
