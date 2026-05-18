package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() (x int) {
	defer func() {
		if p := recover(); p != nil {
			x = 1
		}
	}()

	tempZero := 0
	fail := 1 / tempZero
	fmt.Print(fail)
	panic("?")
}
