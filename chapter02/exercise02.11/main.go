package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		randomNumber := rand.Intn(8)

		if randomNumber%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if randomNumber%2 == 0 {
			fmt.Println("Stop")
			break
		}

		fmt.Println(randomNumber)
	}
}
