package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	// these 2 goroutine are completely decoupled from each other
	// the only dependency is the channel.
	go func(wg *sync.WaitGroup, ch chan int) {
		ch <- 42
		wg.Done()
	}(wg, ch)
	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Println(<-ch)
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
