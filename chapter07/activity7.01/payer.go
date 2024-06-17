package main

import "fmt"

type Payer interface {
	Pay() (string, float64)
}

func payDetails(p Payer) {
	fullName, yearPay := p.Pay()
	fmt.Printf("%s got paid %.2f for the year\n", fullName, yearPay)
}
