package chapter05

import (
	"fmt"
)

// Activity 5.01 – calculating the working hours of employees
func Activity01() {
	chris := Developer{
		Individual: Employee{
			Id:        001,
			FirstName: "Chris",
			LastName:  "Carter",
		},
		HourlyRate: 10,
	}

	chris.LogHours(Sunday, 8)
	chris.LogHours(Monday, 10)

	fmt.Printf("Hours worked on Monday: %d\n", chris.WorkWeek[Sunday])
	fmt.Printf("Hours worked on Monday: %d\n", chris.WorkWeek[Monday])
	fmt.Printf("Hours worked this week: %d\n", chris.HoursWorked())
}

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	var totalHours int

	for _, hours := range d.WorkWeek {
		totalHours += hours
	}

	return totalHours
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Activity 5.02 – calculating the payable amount for employees
// based on working hours

func Activity02() {

}
