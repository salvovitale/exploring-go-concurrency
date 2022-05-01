package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) // bidirectional channel
	wg.Add(2)

	// you do not want that goroutine can be both sending and receiving but in general they should only receive or send.
	// thats why is a good practice to mark them as sending only or receiving only.
	go func(wg *sync.WaitGroup, ch chan int) {
		ch <- 42
		// i m putting this here to avoid that i send and receive from the same goroutine
		time.Sleep(10 * time.Millisecond)
		fmt.Println(<-ch)
		wg.Done()
	}(wg, ch)
	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Println(<-ch)
		ch <- 43
		wg.Done()
	}(wg, ch)

	wg.Wait()
}