package main

import "fmt"

// -- Package level variables
// Declaration without assignment
var mainInt int

// Declaration with assignment
var num = 0

// Visibility in go is determined by first character capitalization
var privateVar = 1 // Lowercase first letter makes a member private to a package
var PublicVar = 2  // Uppercase first letter makes a member public outside of a package

func main() {

	// Declaration within a function
	var localString string
	var localInt int
	var localChar rune // Characters are called runes in go
	var localFloat float64
	values("Zero", localString, localInt, localChar, localFloat)

	// Declaration with assignment within a function
	localString2 := "localString"
	localInt2 := 1
	localChar2 := 'c'
	localFloat2 := 3.14
	values("Set", localString2, localInt2, localChar2, localFloat2)
}

func values(label, s string, i int, c rune, f float64) {
	fmt.Printf("%s Values\n", label)
	fmt.Println("==================")
	fmt.Printf("%T: %v\n%T: %v\n%T: %v\n%T: %v\n", s, s, i, i, c, c, f, f)
	fmt.Println("==================")
}
