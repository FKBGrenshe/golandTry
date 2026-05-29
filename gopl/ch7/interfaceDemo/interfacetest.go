package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	w = new(bytes.Buffer)
	fmt.Println(w)
}
