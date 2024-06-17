package chapter04

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func defineArray() [10]int {
	var arr [10]int
	return arr
}

// Exercise 4.01 – Defining an array
func Exercise01() {
	fmt.Printf("%#v\n", defineArray())
}

// ------------------------------------------------------------------------------------------------

func compArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [5]int{0, 0, 0, 0, 9}

	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}

// Exercise 4.02 – Comparing arrays
func Exercise02() {
	comp1, comp2, comp3 := compArrays()
	fmt.Println("[5]int == [5]int{0}:", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0}:", comp2)
	fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9} :", comp3)
}

// ------------------------------------------------------------------------------------------------

func compArrays02() (bool, bool, [10]int) {
	var arr1 [10]int
	// set key 9 to value 0
	arr2 := [...]int{9: 0}
	// set key 0 to value 1, set key 9 to value 10,
	// and set key 4 to value 5
	arr3 := [10]int{1, 9: 10, 4: 5}

	return arr1 == arr2, arr1 == arr3, arr3
}

// Exercise 4.03 – Initializing an array using keys
func Exercise03() {
	comp1, comp2, arr3 := compArrays02()
	fmt.Println("[10]int == [...]{9:0}:", comp1)
	fmt.Println("[10]int == [10]int{1, 9: 10, 4: 5}}:", comp2)
	fmt.Println("arr3				:", arr3)
}

// ------------------------------------------------------------------------------------------------

func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}

	return fmt.Sprintf("%s %s %s %s", arr[1], arr[0], arr[3], arr[2])
}

// Exercise 4.04 – Reading a single item from an array
func Exercise04() {
	fmt.Println(message())
}

// ------------------------------------------------------------------------------------------------

func message02() string {
	arr := [4]string{"ready", "Get", "Go", "to"}

	arr[1] = "It’s"
	arr[0] = "time"

	return fmt.Sprintf("%s %s %s %s", arr[1], arr[0], arr[3], arr[2])
}

// Exercise 4.05 – Writing to an array
func Exercise05() {
	fmt.Println(message02())
}

// ------------------------------------------------------------------------------------------------

func message03() string {
	m := ""
	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}

	return m
}

// Exercise 4.06 – Looping over an array using a “for i” loop
func Exercise06() {
	fmt.Print(message03())
}

// ------------------------------------------------------------------------------------------------

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	return arr
}

func opArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}

	return arr
}

// Exercise 4.07 – Modifying the contents of an array in a loop
func Exercise07() {
	var arr [10]int
	arr = fillArray(arr)
	arr = opArray(arr)
	fmt.Println(arr)
}

// ------------------------------------------------------------------------------------------------

func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(0)
	}

	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func findLogest(args []string) string {
	var longest string

	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}

	return longest
}

// Exercise 4.08 – Working with slices
func Exercise08() {
	if longest := findLogest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}
}

// ------------------------------------------------------------------------------------------------

func getPassedArgs02() []string {
	var args []string

	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func getLocales(extraLocales []string) []string {
	var locales []string

	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocales...)

	return locales
}

// Exercise 4.09 – Appending multiple items to a slice
func Exercise09() {
	locales := getLocales(getPassedArgs02())
	fmt.Println("Locales to use: ", locales)
}

// ------------------------------------------------------------------------------------------------

func message04() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := fmt.Sprintln("First :", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("Last:", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	m += fmt.Sprintln("First 5 :", s[:5])
	m += fmt.Sprintln("Last 4 :", s[5:])
	m += fmt.Sprintln("Middle 5:", s[2:7])

	return m
}

// Exercise 4.10 – Creating slices from a slice
func Exercise10() {
	fmt.Print(message04())
}

// ------------------------------------------------------------------------------------------------

func genSlices() ([]int, []int, []int) {
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)

	return s1, s2, s3
}

// Exercise 4.11 – Using make to control the capacity of a slice
func Exercise11() {
	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
}

// ------------------------------------------------------------------------------------------------

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99

	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1

	s1 = append(s1, 6)

	s1[3] = 99

	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1

	s1 = append(s1, 6)
	s1[3] = 99

	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

// Exercise 4.12 – Controlling internal slice behavior
func Exercise12() {
	l1, l2, l3 := linked()
	fmt.Println("Linked:", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link:", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link: ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)
}

// ------------------------------------------------------------------------------------------------

func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
	}

	users["073"] = "Tracy"

	return users
}

// Exercise 4.13 – Creating, reading, and writing a map
func Exercise13() {
	fmt.Println("Users:", getUsers())
}

// ------------------------------------------------------------------------------------------------

func getUser(id string) (string, bool) {
	users := getUsers()

	user, exists := users[id]

	return user, exists
}

func Exercise14() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userID := os.Args[1]
	name, exists := getUser(userID)

	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", userID)
		for key, value := range getUsers() {
			fmt.Println(" ID:", key, "Name:", value)
		}
		os.Exit(1)
	}

	fmt.Println("Name:", name)
}

// ------------------------------------------------------------------------------------------------

var usersTwo = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func deleteUser(id string) {
	delete(usersTwo, id)
}

// Exercise 4.15 – Deleting an element from a map
func Exercise15() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userID := os.Args[1]
	deleteUser(userID)
	fmt.Println("Users:", usersTwo)
}

// ------------------------------------------------------------------------------------------------

type id string

func getIDs() (id, id, id) {
	var id1 id
	var id2 id = "1234-5678"
	var id3 id
	id3 = "1234-5678"

	return id1, id2, id3
}

// Exercise 4.16 – Creating a simple custom type
func Exercise16() {
	id1, id2, id3 := getIDs()
	fmt.Println("id1 == id2	  :", id1 == id2)
	fmt.Println("id2 == id3	  :", id2 == id3)

	fmt.Println("id2 == \"1234-5678\":", string(id2) == "1234-5678")
}

// ------------------------------------------------------------------------------------------------

type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

func getUsers02() []user {
	u1 := user{
		name:    "Tracy",
		age:     51,
		balance: 98.43,
		member:  true,
	}

	u2 := user{
		age:  19,
		name: "Nick",
	}

	u3 := user{
		name:    "Bob",
		age:     25,
		balance: 0,
		member:  false,
	}

	u4 := user{}
	u4.name = "Sue"
	u4.age = 31
	u4.member = true
	u4.balance = 17.09

	return []user{u1, u2, u3, u4}
}

// Exercise 4.17 – Creating struct types and values
func Exercise17() {
	users := getUsers02()
	for i, user := range users {
		fmt.Printf("%v: %#v\n", i, user)
	}
}

// ------------------------------------------------------------------------------------------------

type point struct {
	x int
	y int
}

func compare() (bool, bool) {
	point1 := struct {
		x int
		y int
	}{
		10,
		10,
	}

	point2 := struct {
		x int
		y int
	}{}

	point2.x = 10
	point2.y = 5

	point3 := point{10, 10}

	return point1 == point2, point1 == point3
}

// Exercise 4.18 – Comparing structs to each other
func Exercise18() {
	a, b := compare()
	fmt.Println("point == point2:", a)
	fmt.Println("point == point3:", b)
}

// ------------------------------------------------------------------------------------------------

type name string

type location struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

type dot struct {
	name
	location
	size
}

func getDots() []dot {
	var dot1 dot

	dot2 := dot{}
	dot2.name = "A"
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	dot3 := dot{
		name: "B",
		location: location{
			x: 13,
			y: 27,
		},
		size: size{
			width:  5,
			height: 7,
		},
	}

	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.height = 43

	return []dot{dot1, dot2, dot3, dot4}
}

// Exercise 4.19 – Struct embedding and initialization
func Exercise19() {
	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i, dots[i])
	}
}

// ------------------------------------------------------------------------------------------------

func convert() string {
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14

	m := fmt.Sprintf("int8 = %v > int64 = %v\n", i8, int64(i8))
	m += fmt.Sprintf("int = %v > int8 = %v\n", i, int8(i))
	m += fmt.Sprintf("int8 = %v > float32 = %v\n", i8, float64(i8))
	m += fmt.Sprintf("float64 = %v > int = %v\n", f64, int(f64))

	return m
}

// Exercise 4.20 – Numeric type conversions
func Exercise20() {
	fmt.Println(convert())
}

// ------------------------------------------------------------------------------------------------

func doubler(v interface{}) (string, error) {

	switch t := v.(type) {
	case string:
		return t + t, nil
	case bool:
		if t {
			return "truetrue", nil
		}
		return "falsefalse", nil
	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}
		return fmt.Sprint(t.(float32) * 2), nil
	case int:
		return fmt.Sprint(t * 2), nil
	case int8:
		return fmt.Sprint(t * 2), nil
	case int16:
		return fmt.Sprint(t * 2), nil
	case int32:
		return fmt.Sprint(t * 2), nil
	case int64:
		return fmt.Sprint(t * 2), nil
	case uint:
		return fmt.Sprint(t * 2), nil
	case uint8:
		return fmt.Sprint(t * 2), nil
	case uint16:
		return fmt.Sprint(t * 2), nil
	case uint32:
		return fmt.Sprint(t * 2), nil
	case uint64:
		return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("unsupported type passed")
	}
}

// Exercise 4.21 – Type assertions
func Exercise21() {
	res, _ := doubler(-5)
	fmt.Println("-5  :", res)

	res, _ = doubler(5)
	fmt.Println("5   :", res)

	res, _ = doubler("yum")
	fmt.Println("yum :", res)

	res, _ = doubler(true)
	fmt.Println("true:", res)

	res, _ = doubler(float32(3.14))
	fmt.Println("3.14:", res)
}
