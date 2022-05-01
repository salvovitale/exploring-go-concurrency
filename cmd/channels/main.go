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
		wg.Done()
	}(wg, ch)
	go func(wg *sync.WaitGroup, ch <-chan int) {
		fmt.Println(<-ch)
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
