package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var balance int

func ReadBalance() int {
	mu.RLock()         // readers lock
	defer mu.RUnlock() // readers unlock

	time.Sleep(1 * time.Second)
	return balance
}

func main() {
	go func() {
		fmt.Println("read 1")
		fmt.Println(ReadBalance())
	}()

	go func() {
		fmt.Println("read 2")
		fmt.Println(ReadBalance())
	}()

	go func() {
		fmt.Println("read 3")
		fmt.Println(ReadBalance())
	}()
	time.Sleep(5 * time.Second)
}
