package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var tickts =100
var wg sync.WaitGroup
var mu sync.Mutex
func main() {
	wg.Add(4)
	go saleTickts("售票窗口1")
	go saleTickts("售票窗口2")
	go saleTickts("售票窗口3")
	go saleTickts("售票窗口4")
	wg.Wait()
	fmt.Println("所以都卖完了...")
}

func saleTickts(name string)  {
	rand.Seed(time.Now().UnixNano())
	for{
		mu.Lock()
		if tickts>0{
			time.Sleep(time.Duration(rand.Intn(1000)))
			fmt.Println(name,":",tickts)
			tickts--
		}else{
			fmt.Println(name,",买票结束了")
			mu.Unlock()
			break
		}
		mu.Unlock()
	}
	wg.Done()
}
