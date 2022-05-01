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
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2) // we should add 1 for each goroutine
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		// wg.Add(1) we add 2 above. Thats why we do not need to add 1 here
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				cache[id] = b
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		// we cant remove this time sleep here otherwise we will generate a race condition as both cache and query to database are trying to accessing the cache.
		time.Sleep(150 * time.Millisecond)
	}
	// now we can replace the time sleep with the proper tool
	wg.Wait()
}

func queryCache(id int) (database.Book, bool) {
	// we are reading here and we are writing at line 31 this can create a race condition.
	b, ok := cache[id]
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
