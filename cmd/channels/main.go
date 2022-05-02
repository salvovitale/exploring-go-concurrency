package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) // bidirectional channel
	wg.Add(2)

	// you do not want that goroutine can be both sending and receiving but in general they should only receive or send.
	// thats why is a good practice to mark them as sending only or receiving only.
	go func(wg *sync.WaitGroup, ch chan<- int) {
		ch <- 42
		// if i close the channel and try to send again this will panic
		// close(ch)
		// ch <- 27
		wg.Done()
	}(wg, ch)
	// we need to update the channel because only sending type can be closed.
	// closing a channel always has to be on the sending side of the operation.
	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Println(<-ch)
		close(ch)
		// in this case i will get a 0
		fmt.Println(<-ch)
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
