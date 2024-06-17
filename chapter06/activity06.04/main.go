// Activity 6.04 – preventing a panic from crashing the app
package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d directDeposit) validateRoutingNumber() {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrInvalidRoutingNumber {
				fmt.Println(ErrInvalidRoutingNumber)
			}
		}
	}()

	panic(ErrInvalidRoutingNumber)
}

func (d directDeposit) validateLastName() error {
	if !(len(d.lastName) == 0) {
		return nil
	}

	return ErrInvalidLastName
}

func (d directDeposit) report() {
	fmt.Println("Last Name:", d.lastName)
	fmt.Println("First Name:", d.firstName)
	fmt.Println("Bank Name:", d.bankName)
	fmt.Println("Routing Number:", d.routingNumber)
	fmt.Println("Account Number:", d.accountNumber)
}

func main() {
	abe := directDeposit{
		firstName:     "Abe",
		bankName:      "WilkesBooth Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	abe.validateRoutingNumber()
	fmt.Println(abe.validateLastName())
	fmt.Println("**********************************************************************")
	abe.report()
}
