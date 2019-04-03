package main

import "fmt"

type Interface interface {
	UpdateInt(int)
	UpdateString(str string)
	GetValues() (int, string)
}

type impl struct {
	IntValue    int
	StringValue string
}

func (i *impl) UpdateInt(num int) {
	i.IntValue = num
}

func (i *impl) UpdateString(s string) {
	i.StringValue = s
}

func (i *impl) GetValues() (int, string) {
	return i.IntValue, i.StringValue
}

func main() {
	var inf Interface
	inf = &impl{}
	inf.UpdateInt(1)
	inf.UpdateString("string")
	f := inf.GetValues
	fmt.Println(f())

}
