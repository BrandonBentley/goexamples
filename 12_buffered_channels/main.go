package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	buffered := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// producer
	go func() {
		for i := 0; i < 12; i++ {
			buffered <- i
			fmt.Printf("Inserted %v into buffered channel\n", i)
		}
		close(buffered)
		fmt.Println("\nBuffered channel has been closed by the sender\n")
		wg.Done()
	}()
	// consumer
	go func() {
		time.Sleep(time.Second * 2)
		for {
			value, open := <-buffered
			if !open {
				wg.Done()
				break
			}
			fmt.Printf("Read %v from buffered channel\n", value)
			time.Sleep(time.Millisecond * 100)
		}
	}()
	wg.Wait()
}
