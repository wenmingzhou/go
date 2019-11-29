package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("main 即将锁定mutex")
	mutex.Lock()
	fmt.Println("main 已经锁定mutex")
	for i:=1;i<=3;i++ {
		go func(i int) {
			fmt.Println("子goroutime",i,"即将锁定mutex..")
			mutex.Lock()
			fmt.Println("子goroutime",i,"已经锁定mutex..")
		}(i)
	}
	time.Sleep(5*time.Second)
	fmt.Println("main 即将解锁")
	mutex.Unlock()
	fmt.Println("main 已经解锁")
	time.Sleep(3*time.Second)

}
