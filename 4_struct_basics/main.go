package main

import "fmt"

type NewType struct {
	Field1 int
	Field2 string
}

func main() {
	newType := NewType{
		Field1: 1,
		Field2: "",
	}
	fmt.Println(newType)

	var newType2 NewType
	fmt.Println(newType2)
}
