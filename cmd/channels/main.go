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
	// to solve the problem we can use the range operator on the channel in the for loop.
	// and close the channel after we are done sending the messages.
	go func(wg *sync.WaitGroup, ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(wg, ch)
	go func(wg *sync.WaitGroup, ch <-chan int) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
