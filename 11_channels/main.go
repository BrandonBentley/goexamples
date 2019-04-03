package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unbuffered := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// producer
	go func() {
		for i := 0; i < 10; i++ {
			unbuffered <- i
			time.Sleep(time.Millisecond * 100)
		}
		wg.Done()
	}()
	// consumer
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-unbuffered)
		}
		wg.Done()
	}()
	wg.Wait()
}
