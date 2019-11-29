package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond :=sync.Cond{L:&mutex}
	condition :=false

	go func() {
		time.Sleep(1*time.Second)
		cond.L.Lock()
		fmt.Println("子goroutine更改条件")
		condition =true
		cond.Signal()
		fmt.Println("子goroutine解锁.....")
		cond.L.Unlock()
	}()

	cond.L.Lock()
	fmt.Println("main ...已经锁定........")
	if !condition{
		fmt.Println("main ...即将等待........")
		cond.Wait()
		fmt.Println("main ...被唤醒........")
	}

	fmt.Println("main ...继续........")
	fmt.Println("main ...解锁........")
	cond.L.Unlock()


}
