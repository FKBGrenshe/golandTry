package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x < 3; x++ {
			fmt.Printf("counter: put %d into naturals\n", x)
			naturals <- x
		}
		close(naturals)
		return
	}()

	//	squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				close(squares)
				return
			}
			fmt.Printf("squarer: receive %d from naturals\n", x)
			x2 := x * x
			fmt.Printf("squarer: put %d into squares\n", x2)
			squares <- x2
		}
	}()

	// printer (in main goroutine)
	for {
		finalVal, ok := <-squares
		if !ok {
			return
		}
		fmt.Printf("printer: receive %d from squares\n", finalVal)
	}
}
