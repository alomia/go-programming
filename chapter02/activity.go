package chapter02

import (
	"fmt"
	"strconv"
)

/*
Activity 2.01 – looping over map data using range

	words := map[string]int{
			"Gonna": 3,
			"You":   3,
			"Give":  2,
			"Never": 1,
			"Up":    4,
		}
*/
func Activity01(words map[string]int) {
	var popularWord string
	var countWord int

	for word, total := range words {
		if total > countWord {
			popularWord = word
			countWord = total
		}
	}

	fmt.Println("Most popular word:", popularWord)
	fmt.Println("With a count of:", countWord)
}

// Activity 2.02 – implementing FizzBuzz
func Activity02(input int) {
	var output string

	if input%3 == 0 {
		output = "Fizz"
	}

	if input%5 == 0 {
		output += "Buzz"
	}

	if output == "" {
		output = strconv.Itoa(input)
	}

	fmt.Println(output)
}

// Activity 2.03 – bubble sort
func Activity03(numbers []int) {
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(numbers); i++ {
			if numbers[i-1] > numbers[i] {
				numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
				swapped = true
			}
		}
	}
}
