package main

import "fmt"

func main() {
	months := [...]string{
		1:  "january",
		2:  "february",
		3:  "march",
		4:  "april",
		5:  "may",
		6:  "june",
		7:  "july",
		8:  "august",
		9:  "september",
		10: "october",
		11: "november",
		12: "december",
	}
	fmt.Println(months)

	q2 := months[4:7]
	fmt.Println(q2)
	summer := months[6:9]
	fmt.Println(summer)

	fmt.Println(len(q2))
	fmt.Println(cap(q2))
	fmt.Println(len(summer))
	fmt.Println(cap(summer))

	newslice := q2[:6]
	fmt.Println(newslice)

	org := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(org)
	reverseSlice(org)
	fmt.Println(org)

	fmt.Println("一个零值的slice等于nil。一个nil值的slice并没有底层数组")
	var s []int
	fmt.Println(s)
	fmt.Println(s == nil)

	fmt.Println("nil slice 的长度和容量都为0")
	var nilSlice []int
	fmt.Println(len(nilSlice))
	fmt.Println(cap(nilSlice))
	fmt.Println(" 空slice 的长度和容量都为0")
	var emptySlice = []int{}
	fmt.Println(len(emptySlice))
	fmt.Println(cap(emptySlice))

	appendTest()

	nonemptystrings := nonempty([]string{"1", " ", " ", "2"})
	fmt.Println(nonemptystrings)
}

func reverseSlice(org []int) {
	for i, j := 0, len(org)-1; i < j; i, j = i+1, j-1 {
		org[i], org[j] = org[j], org[i]
	}
}

func equals(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendTest() {
	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != " " && s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// hel lo
// hello
