package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "worker1")
	go worker(ctx, "worker2")
	time.Sleep(1 * time.Second)
	cancel() // 发送取消信号，两个worker 都会退出
	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "ctx.Done")
			return
		default:
			fmt.Println(name, "working")
			time.Sleep(time.Second)
		}
	}
}
