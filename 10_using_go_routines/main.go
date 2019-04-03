package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go PrintOutMessageAfterNMilliSeconds("Routine 1 printed", 967, wg)
	go PrintOutMessageAfterNMilliSeconds("Routine 2 printed", 745, wg)
	go PrintOutMessageAfterNMilliSeconds("Routine 3 printed", 555, wg)
	wg.Wait()
	fmt.Println("Exited Successfully")
}

func PrintOutMessageAfterNMilliSeconds(message string, milliseconds int, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * time.Duration(milliseconds))
		fmt.Println(message)
	}
	wg.Done()
}
