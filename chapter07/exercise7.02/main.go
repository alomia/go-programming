// Exercise 7.02 – calculating the area of different shapes using polymorphism
package main

func main() {
	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{length: 20, width: 10}
	s := square{side: 10}

	printShapeDetails(t, r, s)
}
