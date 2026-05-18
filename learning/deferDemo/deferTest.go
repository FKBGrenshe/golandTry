package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("this is defer 1")
	defer fmt.Println("this is defer 2")

	{
		defer fmt.Println("this is defer 3")
		fmt.Println("block ends")
	}

	fmt.Println("this is main")

	fmt.Println("timeDefer start")
	timeDefer()

}

func timeDefer() {
	startAt := time.Now()
	defer func() {
		fmt.Println(time.Since(startAt))
	}()

	time.Sleep(2 * time.Second)
}
