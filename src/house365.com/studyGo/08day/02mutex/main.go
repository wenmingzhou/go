package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁 保证同一时间只有一个goroutine进入临界区，其他的都在等待
func add() {
	defer wg.Done()

	for i := 0; i < 5000000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}

}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
