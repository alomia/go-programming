package chapter03

import (
	"fmt"
	"math"
	"math/big"
	"unicode"
)

// Exercise 3.01 – Program to measure password complexity
func Exercise01(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, ch := range pwR {
		if unicode.IsUpper(ch) {
			hasUpper = true
		}

		if unicode.IsLower(ch) {
			hasLower = true
		}

		if unicode.IsNumber(ch) {
			hasNumber = true
		}

		if unicode.IsPunct(ch) || unicode.IsSymbol(ch) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

// Exercise 3.02 – Floating-point number accuracy
func Exercise02() {
	var a int = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println(a / 3)
	fmt.Println(b / 3)
	fmt.Println(c / 3)
}

// Exercise 3.03 – Triggering number wraparound
func Exercise03() {
	var a int8 = 125
	var b uint8 = 253

	for i := 0; i < 5; i++ {
		a++
		b++

		fmt.Println(i, ")", "int8 ", a, "uint8 ", b)
	}
}

// Exercise 3.04 – Big numbers
func Exercise04() {
	intA := math.MaxInt64
	intA = intA + 1

	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int	:", intA)
	fmt.Println("Big Int : ", bigA.String())
}

// Exercise 3.05 – Safely looping over a string
func Exercise05() {
	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}
