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
		if b, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			cache[id] = b
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book not found id: '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}
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
