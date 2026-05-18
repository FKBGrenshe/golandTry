package main

func outer() func() {
	// 外层函数的变量
	count := 0

	// 匿名函数（闭包）访问外层函数的变量
	return func() {
		count++
		println("Count:", count)
	}
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
