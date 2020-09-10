package main

import "fmt"

func main() {
	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	sundayIndex := len(days) - 1
	days = append(days[sundayIndex:], days[:sundayIndex]...)
	fmt.Println(days)
}
