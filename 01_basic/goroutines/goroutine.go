package main

import (
	"fmt"
	"time"
)

// goroutine function
func todo(id int, ch chan string) {
	// logic
	time.Sleep(time.Millisecond * 300)

	// send a msg to channel
	ch <- fmt.Sprintf("todo %d done!", id)
}

func main() {

	// create a channel string
	// channel are used to communicate to goroutine
	ch := make(chan string)

	// started 5 goroutines
	// everyone will run todo() in concurrency
	for i := 1; i <= 5; i++ {
		go todo(i, ch) // 'go' create a new goroutine
	}

	// main wait 5 msg from todo
	// this grated sync: it will continue when every todo send something
	for i := 1; i <= 5; i++ {
		msg := <-ch // recive from channel (block until recive)
		fmt.Println(msg)
	}

	fmt.Println("All goroutine done!")
}