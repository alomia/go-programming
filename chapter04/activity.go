package chapter04

import (
	"fmt"
	"os"
	"strings"
)

// Activity 4.01 – Filling an array
func Activity01() {
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	fmt.Println(arr)
}

// ------------------------------------------------------------------------------------------------

var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func getUser02(id string) (string, bool) {
	user, exists := users[id]

	return user, exists
}

// Activity 4.02 – Printing a user’s name based on user input
func Activity02() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userID := os.Args[1]
	user, exists := getUser02(userID)

	if !exists {
		fmt.Printf("The user with id %s does not exist.\n", userID)
		os.Exit(1)
	}

	fmt.Println("Hi", user)
}

// ------------------------------------------------------------------------------------------------

// Activity 4.03 – Slicing the week
func Activity03() {
	weekdays := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	weekdays = append(weekdays[6:], weekdays[:6]...)

	fmt.Println(weekdays)
}

// ------------------------------------------------------------------------------------------------

// Activity 4.04 – Removing an element from a slice
func Activity04() {
	sl := []string{"Good", "Good", "Bad", "Good", "Good"}
	sl = append(sl[:2], sl[3:]...)

	fmt.Println(sl)
}

// ------------------------------------------------------------------------------------------------

// Activity 4.05 – Creating a locale checker

type locale struct {
	language  string
	territory string
}

func getLocales02() map[locale]struct{} {
	supportedLocales := make(map[locale]struct{}, 5)
	supportedLocales[locale{"en", "US"}] = struct{}{}
	supportedLocales[locale{"en", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "FR"}] = struct{}{}
	supportedLocales[locale{"ru", "RU"}] = struct{}{}
	return supportedLocales
}

func localeExists(l locale) bool {
	_, exists := getLocales02()[l]
	return exists
}

func Activity05() {
	if len(os.Args) < 2 {
		fmt.Println("No locale passed")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	passedLocale := locale{
		territory: localeParts[1],
		language:  localeParts[0],
	}

	if !localeExists(passedLocale) {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("Locale passed is supported")
}

// ------------------------------------------------------------------------------------------------

func typeChecker(v any) string {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

// Activity 4.06 – Type checker
func Activity06() {
	dataTypes := []any{
		int32(1),
		float64(3.14),
		"hello",
		true,
		struct{}{},
	}

	for _, v := range dataTypes {
		fmt.Println(v, "is", typeChecker(v))
	}
}
