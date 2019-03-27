package main

import "fmt"

type AttachToMe struct {
	Field1 int
	Field2 string
}

func (o AttachToMe) TryUpdateField1(f1 int) {
	fmt.Println("Attempt to set Field1 to", f1)
	o.Field1 = f1
}

func (o AttachToMe) GetField1() int {
	return o.Field1
}

func main() {
	attachToMe := AttachToMe{
		Field1: 1,
		Field2: "",
	}
	fmt.Println("attachToMe struct: ", attachToMe)
	fmt.Println("attachToMe.Field1: ", attachToMe.Field1)
	attachToMe.TryUpdateField1(4)
	fmt.Println("attachToMe.Field1: ", attachToMe.Field1)
	pointerToAttachToMe := &attachToMe
	pointerToAttachToMe.UpdateField1(4)
	fmt.Println("attachToMe.Field1: ", pointerToAttachToMe.Field1)
}

func (o *AttachToMe) UpdateField1(f1 int) {
	fmt.Println("Attempt to set Field1 to", f1)
	o.Field1 = f1
}
