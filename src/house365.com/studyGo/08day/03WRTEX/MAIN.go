package main

import (
	"fmt"
	"sync"
	"time"
)

//读锁和写锁。读写互斥锁
// 当一个goroutine获取读锁之后， 其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
// 当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待

var x = 0
var lock sync.Mutex
var rwlock sync.RWMutex
var wg sync.WaitGroup

func read() {
	defer wg.Done()
	lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	lock.Unlock()
}
func write() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}

	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
