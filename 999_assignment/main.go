package main

import (
	"os"
	"time"

	"github.com/BrandonBentley/goexamples/999_assignment/factory"
)

func main() {
	fact := factory.New()
	fact.Add("first")
	fact.Add("second")
	fact.Add("third")
	fact.Add("fourth")
	fact.StartUp()
	for _, v := range os.Args[1:] {
		fact.Add(v)
		time.Sleep(time.Millisecond * 100)
	}
	fact.Shutdown()
}
