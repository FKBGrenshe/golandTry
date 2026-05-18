package main

import (
	"fmt"
)

func main() {
	defer func() {
		if info := recover(); info != nil {
			fmt.Print("this is recover!")
			fmt.Println(info)
		}
	}()

	var val int = 0
	fmt.Scan(val)
	fmt.Println(3 / val)
}
