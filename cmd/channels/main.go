package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) // bidirectional channel
	wg.Add(2)

	// as we know the sender and the receiver need to send receive the same number of messages.
	// otherwise we are generating a deadlock.
	go func(wg *sync.WaitGroup, ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}(wg, ch)
	go func(wg *sync.WaitGroup, ch <-chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
