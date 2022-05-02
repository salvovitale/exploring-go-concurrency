package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) // bidirectional channel
	wg.Add(2)

	go func(wg *sync.WaitGroup, ch chan<- int) {
		// ch <- 42
		close(ch)
		wg.Done()
	}(wg, ch)
	go func(wg *sync.WaitGroup, ch <-chan int) {
		if msg, isChannelOpen := <-ch; isChannelOpen {
			fmt.Println(msg)
		} else {
			fmt.Println("Channel is closed!")
		}
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
