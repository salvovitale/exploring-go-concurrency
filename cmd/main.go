package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/salvovitale/exploring-go-concurrency/internal/database"
)

var cache = map[int]database.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id)

		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				cache[id] = b
				fmt.Println(b)
			}
		}(id)
		// fmt.Printf("Book not found id: '%v'", id)
		// if we remove the following sleep call below, nothing will be printed because
		// the go-routine will be schedule but the main thread will be done before they could do something.
		time.Sleep(150 * time.Millisecond)
	}
	// with this following time sleep we will ge the complete output, given the fact that we are
	// in control of the actual time spent in the underlying operation.
	// However using sleeps call does not bring
	// us a long way. There should be a better way of managing go-routines.
	time.Sleep(2 * time.Second)
}

func queryCache(id int) (database.Book, bool) {
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
