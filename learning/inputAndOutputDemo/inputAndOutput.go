package main

import "fmt"

func main() {
	var name string
	name = "cc"
	fmt.Println(name)

	var userName string = "CC"
	fmt.Println(userName)

	othername := "c"
	fmt.Println(othername)

	fmt.Println("枫枫知道")
	fmt.Println(1)
	fmt.Println(true)
	fmt.Println("什么", "都", "可以", "输出")

	name = fmt.Sprintf("%v", "你好")
	fmt.Println(name)

	fmt.Println("输入您的名字：")
	var cname string
	fmt.Scan(&cname) // 这里记住，要在变量的前面加个&, 后面讲指针会提到
	fmt.Println("你输入的名字是", cname)

	var a1 int
	var a2 float32
	var a3 string
	var a4 bool

	fmt.Printf("%#v\n", a1)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)
	fmt.Printf("%#v\n", a4)

	// call func
	sayHello()

	var sList = []string{"a", "b", "c", "d"}
	fmt.Println(sList[len(sList)]) // 拿到a这个元素

}

func sayHello() {
	fmt.Println("hello")
}

func add(numList ...int) int {

	sum := 0

	for i := range numList {
		sum += numList[i]
	}

	fmt.Println(sum)

	return sum
}

func tryAnoymusFunc() {
	var temp = func(a, b int) int {
		return a + b
	}
	fmt.Println(temp)
}
