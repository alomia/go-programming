package chapter05

import (
	"fmt"
	"strings"
)

// Exercise 5.01 – creating a function to print salesperson
// expectation ratings from the number of items sold
func Exercise01() {
	itemsSold()
}

func itemsSold() {
	items := make(map[string]int)
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)

		if v < 40 {
			fmt.Println("is below expectations.")
		} else if v > 40 && v < 100 {
			fmt.Println("meets expectations.")
		} else if v > 100 {
			fmt.Println("exceeded expectations.")
		}
	}
}

// Exercise 5.02 – mapping index values to column headers
func Exercise02() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)

	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	csvHdrCol(hdr2)
}

func csvHdrCol(header []string) {
	csvHeadersToColumnIndex := make(map[int]string)

	for i, v := range header {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}

	fmt.Println(csvHeadersToColumnIndex)
}

// Exercise 5.03 – creating a checkNumbers function with return values
func Exercise03() {
	for i := 0; i <= 15; i++ {
		num, result := checkNumbers(i)
		fmt.Printf("Results: %d %s\n", num, result)
	}
}

func checkNumbers(i int) (int, string) {
	switch {
	case i%2 == 0:
		return i, "Even"
	default:
		return i, "Odd"
	}
}

// Exercise 5.04 – mapping a CSV index to a column header with return values
func Exercise04() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrCol02(hdr)
	fmt.Println("Result: ")
	fmt.Println(result)
	fmt.Println()

	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	result2 := csvHdrCol02(hdr2)
	fmt.Println("Result2: ")
	fmt.Println(result2)
	fmt.Println()
}

func csvHdrCol02(hdr []string) map[int]string {
	csvIdxToCol := make(map[int]string)

	for i, v := range hdr {
		v = strings.TrimSpace(v)

		switch strings.ToLower(v) {
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v

		}
	}

	return csvIdxToCol
}

// Exercise 5.05 – summing numbers
func Exercise05() {
	nums := []int{5, 10, 15}
	fmt.Println("Total: ", sum(8, 3))
	fmt.Println("Total: ", sum(nums...))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Exercise 5.06 – creating an anonymous function to calculate the square root of a number
func Exercise06() {

	j := 9
	x := func(i int) int {
		return i * i
	}
	fmt.Printf("The square of %d is %d\n", j, x(j))

}

// Exercise 5.07 – creating a closure function to decrement a counter
func Exercise07() {
	counter := 4
	x := decrement(counter)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func decrement(i int) func() int {
	return func() int {
		i--
		return i
	}
}

// Exercise 5.08 – creating various functions to calculate salary
func Exercise08() {
	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)

	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Developer salary: %d\n", devSalary)
}

func salary(x, y int, f func(int, int) int) int {
	pay := f(x, y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}
