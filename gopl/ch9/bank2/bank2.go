package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sema = make(chan struct{}, 1) // a binary semaphore
	mu   sync.Mutex               // a lock

	balance int
)

func DepositByChannel(amount int) {
	sema <- struct{}{} // acquire the semaphore
	balance += amount
	<-sema // release the semaphore
}

func DepositByLock(amount int) {
	mu.Lock()
	defer mu.Unlock()

	balance += amount
}

func BalanceByChannel() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema
	return b
}

func BalanceByLock() int {
	mu.Lock()
	defer mu.Unlock()

	b := balance
	return b
}

func Deposit(amount int) {
	balance += amount
}

func Balance() int {
	return balance
}

func main() {
	fmt.Println("concurrent safety by no one")
	taskPool(Deposit, Balance, 500) // 增加并发数量
	time.Sleep(2 * time.Second)     // 增加等待时间
	fmt.Println("balance: ", Balance())
	balance = 0
	fmt.Println("concurrent safety by channel")
	taskPool(DepositByChannel, BalanceByChannel, 500)
	time.Sleep(2 * time.Second)
	fmt.Println("balance: ", BalanceByChannel())
	balance = 0
	fmt.Println("concurrent safety by lock")
	taskPool(DepositByLock, BalanceByLock, 500)
	time.Sleep(2 * time.Second)
	fmt.Println("balance: ", BalanceByLock())
	balance = 0

}

func taskPool(add func(amount int), get func() int, concurrentNum int) {
	for i := range concurrentNum {
		go func(i int) {
			// 增加一些延迟来制造竞态条件
			time.Sleep(time.Duration(i%10) * time.Millisecond)
			add(10)
		}(i)
	}
}
