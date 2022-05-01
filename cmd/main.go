package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/salvovitale/exploring-go-concurrency/internal/database"
)

var cache = map[int]database.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2) // we should add 1 for each goroutine
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		// wg.Add(1) we add 2 above. Thats why we do not need to add 1 here
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				// this will lock other writers and readers.
				m.Lock()
				cache[id] = b
				m.Unlock()
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
	}
	// now we can replace the time sleep with the proper tool
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (database.Book, bool) {
	// this will allow multiple readers but not writers.
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int) (database.Book, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range database.GetBooks() {
		if b.ID == id {
			return b, true
		}
	}

	return database.Book{}, false
}
