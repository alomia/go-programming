package main

import "fmt"

func message() string {
	var msg string

	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		arr[i] *= arr[i]
		msg += fmt.Sprintf("%v: %v\n", i, arr[i])
	}

	return msg
}

func main() {
	fmt.Print(message())
}
