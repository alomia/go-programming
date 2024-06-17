// Activity 7.01 – calculating pay and performance review
package main

import (
	"fmt"
	"os"
)

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d := Developer{Individual: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := Manager{Individual: Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	payDetails(d)
	payDetails(m)
}
