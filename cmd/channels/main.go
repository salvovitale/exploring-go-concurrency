package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// Channels are blocking by default.
	// In this example channel is blocking until a message is sent.
	// fmt.Println(<-ch)
	// ch <- 42

	// in this case channel is blocking until there is a receiver.
	ch <- 42
	fmt.Println(<-ch)

	// thats why channels work in the context of goroutines.
}
