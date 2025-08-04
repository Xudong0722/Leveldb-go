package main

import (
	"fmt"

	"github.com/Xudong0722/Leveldb-go/db"
	"github.com/Xudong0722/Leveldb-go/utils"
)

// This file is for Debug.

func main() {
	sl := db.NewSkipList(utils.StringComprator)
	sl.Insert(string("test"))
	res1 := sl.Contains(string("test"))
	fmt.Println(res1)
}
