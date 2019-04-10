package main

import (
	"fmt"
	"time"
)

func main() {
	producer := make(chan int, 5)
	consumerA := make(chan int)
	consumerB := make(chan int)
	go Output(consumerA, "ConsumerA")
	go Output(consumerB, "ConsumerB")
	go Producer(producer)
	Fanout(producer, consumerA, consumerB)
}

func Producer(Out chan<- int) {
	for i := 0; i < 15; i++ {
		time.Sleep(time.Millisecond * 100)
		Out <- i
	}
	close(Out)
}

func Fanout(In <-chan int, OutA, OutB chan<- int) {
	for data := range In { // Recieve until closed
		select { // Send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}
	close(OutA)
	close(OutB)
}

func Output(In <-chan int, consumerName string) {
	for {
		if data, ok := <-In; ok {
			fmt.Printf("%v: %v\n", consumerName, data)
			time.Sleep(time.Nanosecond)
		} else {
			return
		}
	}

}
