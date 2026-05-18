package main

import "fmt"

func addLoop(args ...int) int {
	sum := 0
	for index, value := range args {
		sum += value
		fmt.Println(index, value)
	}
	return sum
}

func main() {
	fmt.Println(addLoop(1, 2, 3, 4, 5))
	fmt.Println(addLoop(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
