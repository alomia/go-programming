package main

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (d Manager) Pay() (string, float64) {
	fullName := d.Individual.FullName()
	yearPay := d.Salary + (d.Salary * d.CommissionRate)

	return fullName, yearPay
}
