package main

import (
	"fmt"

	"github.com/Xudong0722/Leveldb-go/db"
)

// This file is for Debug.

func main() {
	tb := db.NewMemTable()
	tb.Put([]byte("test-2"), []byte("value-2"))
	// tb.Put([]byte("test-5"), []byte("value-5"))
	// tb.Put([]byte("test-8"), []byte("value-8"))
	// tb.Put([]byte("test-7"), []byte("value-7"))
	// tb.Put([]byte("test-6"), []byte("value-6"))
	// tb.Put([]byte("test-3"), []byte("value-3"))
	// tb.Put([]byte("test-4"), []byte("value-4"))
	tb.Put([]byte("test-0"), []byte("value-0"))

	val, _ := tb.Get([]byte("test-0"))
	fmt.Println(string(val))
}
