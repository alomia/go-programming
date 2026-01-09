package main

import "fmt"

func removeBad() []string {
	words := []string{"Good", "Good", "Bad", "Good", "Good"}
	words = append(words[:2], words[3:]...)
	return words
}

func main() {

	fmt.Println(removeBad())
}
