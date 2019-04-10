package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 5)
	go func() {
		time.Sleep(time.Millisecond * 100)
		channel <- 1
	}()
	data, more, ok := TryRecieveWithTimeout(channel, time.Second)
	if !ok {
		fmt.Println("channel recieve timed out")
	} else {
		fmt.Println(data, more)
	}
}

// adding the <- to the channel parameter works as
// somewhat of a permissions modifier.
func TryRecieveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	case <-time.After(duration): // time.After() returns a channel that is immediately blocking until after the given duration
		return 0, true, false
	}
}
