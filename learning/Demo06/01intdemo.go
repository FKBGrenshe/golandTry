package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int = 10
	fmt.Printf("num=%v type=%T \n", num, num)

	s1 := `first line
second line
third line
\n
`

	s2 := `other line`

	var flag = strings.Contains(s1, s2)
	fmt.Println(flag)
	var str = "123-456-789"
	var arr = strings.Split(str, "-")
	var str2 = strings.Join(arr, "*")
	fmt.Println(str2)
}
