package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Go Routine")
		wg.Done()
	}()
	wg.Wait()
}
