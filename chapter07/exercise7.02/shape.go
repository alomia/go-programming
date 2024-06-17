package main

import "fmt"

type Shape interface {
	Area() float64
	Name() string
}

func printShapeDetails(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", shape.Name(), shape.Area())
	}
}
