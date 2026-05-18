package main

import (
	"fmt"
	"os"
)

var done = make(chan struct{})

// 轮询退出状态
func cancelled() bool {
	select {
	// 检查done通道是否有数据可读（或已关闭）
	case <-done:
		// 如果done通道已关闭：<-done会立即返回零值，进入当前case
		return true
	// 如果上面的case都不满足，执行default分支
	default:
		// 如果done通道未关闭：select会立即执行default分支
		return false
	}
}

func main() {
	// determine the initial directories
	roots := os.Args[1:]
	if len(roots) == 0 {
		fmt.Println("can not scan the root dir, return")
		return
	}

	go func(){
		os.Stdin.Read()
	}
}
