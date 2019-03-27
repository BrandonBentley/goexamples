package main

import "fmt"

func main() {
	f := func(str string, i int) {
		fmt.Println("String: ", str)
		fmt.Println("Int: ", i)
	}

	f("something", 66)
	returningFunction(7)("Returned Function")
}

func returningFunction(i int) func(str string) {
	return func(str string) {
		fmt.Println("String: ", str)
		fmt.Println("Int: ", i)
	}
}
