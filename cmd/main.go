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

// the current code does not have any logic to decided from where to print if from chache or db.
func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan database.Book)
	dbCh := make(chan database.Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2) // we should add 1 for each goroutine
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- database.Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		// wg.Add(1) we add 2 above. Thats why we do not need to add 1 here
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- database.Book) {
			if b, ok := queryDatabase(id); ok {
				// this will lock other writers and readers.
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, dbCh <-chan database.Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				// i need to read from dbChannel as well because otherwise the dbCh will never be drained completely. And I will get a deadlock error.
				// basically if the result will come before from the cache the db result will be discard.
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		// this time sleep should stay to simulate real call which are not simultaneous.
		time.Sleep(150 * time.Millisecond)
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
	time.Sleep(100 * time.Millisecond)
	for _, b := range database.GetBooks() {
		if b.ID == id {
			return b, true
		}
	}

	return database.Book{}, false
}
