package main

import (
	"fmt"
	"golandTry/gopl/ch8/countdown1"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		// read a singl byte
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	select {

	case <-time.After(10 * time.Second):
		// do nothing
	case <-abort:
		fmt.Println("launch cancelled")
		return
	}
	countdown1.Launch()
}
