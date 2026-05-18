package main

import (
	"fmt"
	"strings"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func main() {
	// 使用 defer + recover 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到 panic: %v\n", r)
			fmt.Println("程序恢复执行，没有崩溃")
		}
	}()

	f := square
	fmt.Println(f(5))

	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	// f = product
	// fmt.Println(f(4, 5))

	var g func(int, int) int

	// 在调用前检查函数是否为 nil
	if g == nil {
		fmt.Println("g is nil，正在初始化...")
		g = product
	}
	fmt.Println(g(6, 7))

	fmt.Println(strings.Map(add1, "HAL-9000"))

	// -------------- anonymous function --------------
	fmt.Println(strings.Map(func(r rune) rune { return r + 2 }, "HAL-9000"))

	//// -------------- 匿名函数访问完整词法环境 --------------
	//inc := outer()
	//inc()
	//inc()
	//inc()
	//
	//// -------------- 匿名函数访问完整词法环境 --------------
	//h := squares()
	//fmt.Println(h)   // 输出 1
	//fmt.Println(h()) // 输出 1
	//fmt.Println(h()) // 输出 4
	//fmt.Println(h()) // 输出 9

	/**
	 * squares是外层函数本身，
	 * squares()是外层函数的返回值也就是匿名函数，被赋值给了h，h就是匿名函数，
	 * h()是匿名函数的返回值是int
	 */

}

func add1(r rune) rune {
	return r + 1
}
