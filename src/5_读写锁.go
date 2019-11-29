package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var n int
var WG sync.WaitGroup
var rwm sync.RWMutex
func main() {
	WG.Add(4)
	for i:=1;i<=2;i++ {
		go read(i)
	}

	for i:=1;i<=2;i++ {
		go write(i)
	}
	WG.Wait()
}
func write(i int)  {
	defer WG.Done()
	rand.Seed(time.Now().UnixNano())
	rwm.Lock()
	fmt.Println("写操作:",i,"即将写入数据..")
	randNum:=rand.Intn(100)+1
	n =randNum
	fmt.Println("写操作:",i,"已经结束，写入了,",randNum)
	rwm.Unlock()

}

func read(i int)  {
	defer WG.Done()
	rwm.RLock()
	fmt.Println("读操作:",i,"即将读取数据..")
	v :=i
	fmt.Println("读操作:",i,"已经结束，读取了数据,",v )
	rwm.RUnlock()
}
