package main

import (
	"fmt"
	"sort"
)

func main() {
	var list []string

	list = append(list, "hello")
	list = append(list, "world")

	fmt.Println(list)
	fmt.Println(len(list))

	// 修改第二个元素
	list[1] = "不知道"
	fmt.Println(list)

	//var list_struct []

	strings := make([]string, 10)
	fmt.Println(strings, len(strings), cap(strings))
	fmt.Println(strings == nil)

	fmt.Println("-------------------------------")

	var list1 = [...]string{"a", "b", "c", "d", "e", "f", "g"}
	slices := list1[:]
	fmt.Println(slices)
	fmt.Println(list1[1:2]) // b

	fmt.Println("-------------------------------")
	var list2 = []int{4, 5, 3, 2, 7}
	fmt.Println("排序前:", list2)
	sort.Ints(list2)
	fmt.Println("升序:", list2)

	sort.Sort(sort.Reverse(sort.IntSlice(list2)))
	fmt.Println("降序:", list2)

	sort.Ints(list2)

}

//tryp student struct{
//	name string
//}

//func (receiver ``) name()  {
//
//}
