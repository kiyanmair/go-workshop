package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}
	var topWord string = ""
	var topCount int = 0
	for word, count := range words {
		if count > topCount {
			topWord = word
			topCount = count
		}
	}
	fmt.Println("Most popular word:", topWord)
	fmt.Println("With a count of:", topCount)
}
