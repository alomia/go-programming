package main

import "fmt"

func main() {

	taxTotal := 0.0

	// Cake
	taxTotal += salesTax(0.99, 7.5)

	// Milk
	taxTotal += salesTax(2.75, 1.5)

	// Butter
	taxTotal += salesTax(0.87, 2)

	fmt.Println("Sales Tax Total:", taxTotal)
}

func salesTax(cost float64, taxRate float64) float64 {
	return cost * taxRate / 100
}
