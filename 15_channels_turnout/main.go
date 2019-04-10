package main

import (
	"fmt"
	"time"
)

func main() {
	producerA := make(chan int, 5)
	producerB := make(chan int, 5)
	consumerA := make(chan int)
	consumerB := make(chan int)
	quit := make(chan int)
	go Output(consumerA, "ConsumerA")
	go Output(consumerB, "ConsumerB")
	go Producer(producerA, 0)
	go Producer(producerB, 20)
	go func() {
		time.Sleep(time.Second * 2)
		close(quit)
	}()
	Turnout(quit, producerA, producerB, consumerA, consumerB)
}

func Turnout(Quit <-chan int, InA, InB, OutA, OutB chan int) {
	for {
		select {
		case data := <-InA:
			select {
			case OutA <- data:
			case OutB <- data:
			}
		case data := <-InB:
			select {
			case OutA <- data:
			case OutB <- data:
			}
		case <-Quit:
			close(InA)
			close(InB)

			Fanout(InA, OutA, OutB) //Flush the remaining data
			Fanout(InB, OutA, OutB)
		}

	}
}

func Producer(Out chan<- int, start int) {
	i := start
	for {
		time.Sleep(time.Millisecond * 200)
		Out <- i
		i++
	}
}

func Fanout(In <-chan int, OutA, OutB chan<- int) {
	for data := range In { // Recieve until closed
		select { // Send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}
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
