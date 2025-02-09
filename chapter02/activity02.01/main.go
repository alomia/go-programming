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

	var highestCount int
	var highestWord string

	for word, count := range words {
		if count > highestCount {
			highestCount = count
			highestWord = word
		}
	}

	fmt.Println("Most popular word:", highestWord)
	fmt.Println("With a count of:", highestCount)
}
