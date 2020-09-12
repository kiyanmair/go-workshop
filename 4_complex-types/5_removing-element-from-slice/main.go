package main

import (
	"errors"
	"fmt"
)

func getBadIndex(badSlice []string) (int, error) {
	for index, value := range badSlice {
		if value == "Bad" {
			return index, nil
		}
	}
	return -1, errors.New("no bad value in provided slice")
}

func deleteIndex(oldSlice []string, index int) []string {
	firstPart := oldSlice[:index]
	lastPart := oldSlice[index+1:]
	return append(firstPart, lastPart...)
}

func main() {
	oldSlice := []string{
		"Bad",
		"Good",
		"Good",
		"Bad",
		"Good",
		"Good",
		"Bad",
	}
	newSlice := append(oldSlice[:0:0], oldSlice...)
	for {
		indexToDelete, err := getBadIndex(newSlice)
		if err != nil {
			break
		}
		newSlice = deleteIndex(newSlice, indexToDelete)
	}
	fmt.Println(newSlice)
}
