package main

import "fmt"

import "crypto/sha256"

func main() {
	var a [3]int
	fmt.Print(a)
	fmt.Print(a[0])
	fmt.Print(len(a))

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	fmt.Printf("-----------------\n")

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	arr1 := [3]int{1, 1, 1}
	arr2 := [3]int{1, 1, 1}
	arr3 := [3]int{12, 1, 1}
	fmt.Println(arr1 == arr2) // "true"
	fmt.Println(arr1 == arr3) // "false"

	sign()
}

func sign() {
	fmt.Print("sign start \n")
	sign1 := sha256.Sum256([]byte("x"))
	sign2 := sha256.Sum256([]byte("x"))
	sign3 := sha256.Sum256([]byte("3"))
	fmt.Printf("%x\n%x\n%x\n", sign1, sign2, sign3)
	fmt.Println(sign1 == sign2)
	fmt.Println(sign1 == sign3)

	zero(&sign3)
	fmt.Printf("%x\n", sign3)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
