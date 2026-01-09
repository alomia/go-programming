package main

import "fmt"

func main() {
	daysWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Saturda", "Friday", "Sunday"}

	newWeek := append(daysWeek[6:], daysWeek[:6]...)

	fmt.Println(newWeek)
}
