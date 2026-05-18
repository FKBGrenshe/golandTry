package main

// counter -> [naturals] -> squarer -> [squares] -> printer

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)

}

func counter(naturals chan<- int) {
	for x := 0; x < 100; x++ {
		naturals <- x
	}
	close(naturals)
}
func squarer(naturlas <-chan int, squares chan<- int) {
	for receiveVal := range naturlas {
		sendVal := receiveVal * receiveVal
		squares <- sendVal
	}
	close(squares)
	// close(naturlas) //invalid operation: cannot close receive-only channel naturlas
}
func printer(squares <-chan int) {
	for receiveVal := range squares {
		println(receiveVal)
	}
}
