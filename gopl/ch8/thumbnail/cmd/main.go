package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golandTry/gopl/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	// input := bufio.NewScanner())
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
