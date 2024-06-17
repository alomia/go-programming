package chapter03

import (
	"errors"
	"fmt"
)

func salesTax(cost float64, taxRate float64) float64 {
	return cost * taxRate
}

// Activity 3.01 – Sales tax calculator
func Activity01() {
	taxTotal := .0

	taxTotal += salesTax(.99, .075)

	taxTotal += salesTax(2.75, .015)

	taxTotal += salesTax(.87, .02)

	fmt.Println("Sales Tax Total: ", taxTotal)
}

// ------------------------------------------------------------------------------------------------

const (
	goodScore      = 450
	lowScoreRatio  = 10
	goodScoreRatio = 20
)

var (
	ErrCreditScore = errors.New("invalid credit score")
	ErrIncome      = errors.New("income invalid")
	ErrLoanAmount  = errors.New("loan ammount invalid")
	ErrLoanTerm    = errors.New("loan term not a multiple of 12")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	interest := 20.0
	if creditScore >= goodScore {
		interest = 15.0
	}

	if creditScore < 1 {
		return ErrCreditScore
	}

	if income < 1 {
		return ErrIncome
	}

	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 10

	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	totalInterest := (payment * loanTerm) - loanAmount

	approved := false

	if income > payment {
		ratio := (payment / income) * 100
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = false
		}
	}

	fmt.Println("Credit Score    :", creditScore)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Amount     :", loanAmount)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate            :", interest)
	fmt.Println("Total Cost      :", totalInterest)
	fmt.Println("Approved        :", approved)
	fmt.Println("")

	return nil
}

// Activity 3.02 – Loan calculator
func Activity02() {
	// Approved
	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Denied
	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	err = checkLoan(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
