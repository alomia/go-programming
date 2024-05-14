package chapter02

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("input can't be a negative number")
	} else if input > 100 {
		return errors.New("input can't be over 100")
	} else if input%7 == 0 {
		return errors.New("input can't be divisible by 7")
	} else {
		return nil
	}
}

// Exercise 2.04 – implementing initial if statements
func Exercise04(input int) {
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is Even")
		return
	} else {
		fmt.Println(input, "is Odd")
	}
}

// Exercise 2.05 – using a switch statement
func Exercise05(dayBorn time.Weekday) {
	switch dayBorn {
	case time.Monday:
		fmt.Println("Monday's child is fair of face")
	case time.Tuesday:
		fmt.Println("Tuesday's child is full of grace")
	case time.Wednesday:
		fmt.Println("Wednesday's child is full of woe")
	case time.Thursday:
		fmt.Println("Thursday's child has far to go")
	case time.Friday:
		fmt.Println("Friday's child is loving and giving")
	case time.Saturday:
		fmt.Println("Saturday's child works hard for a living")
	case time.Sunday:
		fmt.Println("Sunday's child is bonny and blithe")
	default:
		fmt.Println("Error, day born not valid")
	}
}

// Exercise 2.06 – switch statements and multiple case values
func Exercise06(dayBorn time.Weekday) {
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Error, day born not valid")
	}
}

// Exercise 2.07 – expressionless switch statements
func Exercise07(dayBorn time.Weekday) {
	switch {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Born some other day")
	}
}

// Exercise 2.08 – using a for i loop
func Exercise08() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// Exercise 2.09 – looping over arrays and slices
func Exercise09() {
	names := []string{"Jim", "Jane", "Joe", "June"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

// Exercise 2.10 – looping over a map
func Exercise10() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}

// Exercise 2.11 – using break and continue to control loops
func Exercise11() {
	for {
		r := rand.Intn(8)
		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		}

		fmt.Println(r)
	}
}
