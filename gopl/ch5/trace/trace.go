package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s, time at: %v", msg, start)

	return func() {
		end := time.Now()
		timeDuration := end.Sub(start)
		log.Printf(" exits %s, time at %v, time duration = %v", msg, end, timeDuration)
	}
}

func option() {
	time.Sleep(10 * time.Second)
}

func main() {

	defer trace("option")()
	// defer trace("option")() - 分两部分
	// 1. 立即执行部分 trace("option")
	// 		原因：Go语言的defer语句中，所有函数的参数都在声明时立即求值
	// 2. 延迟执行部分 func(){end := time.Now()...} -- 只有这个返回的匿名函数被推迟到main函数返回时执行
	option()

}
