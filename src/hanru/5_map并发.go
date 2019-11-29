package main

import (
	"fmt"
	"sync"
	"time"
)
var mutex  sync.Mutex
func main() {
	map1 :=make(map[int] string)
	for i:=1;i<=10;i++{
		go func(i int) {
			mutex.Lock()
			map1[i] =fmt.Sprint("数据",i)
			mutex.Unlock()
		}(i)
	}
	time.Sleep(2*time.Second)
	fmt.Println(map1)
}
