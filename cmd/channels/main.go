package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // buffered channel
	// un-buffered channel are having 0 buffering so there should always be a matching number of senders and receivers

	wg.Add(2)
	// these 2 goroutine are completely decoupled from each other
	// the only dependency is the channel.
	go func(wg *sync.WaitGroup, ch chan int) {
		ch <- 42
		// if i add another message i will get a deadlock because there is only one receiver.
		// to fix this we need buffered channels.
		ch <- 3
		wg.Done()
	}(wg, ch)
	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Println(<-ch)
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
