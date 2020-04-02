package main

import (
	"fmt"
	"sync"
	"time"
)

var exitChan = make(chan bool)
var wg sync.WaitGroup

func f1() {
	defer wg.Done()
Loop:
	for {
		fmt.Println("周文明")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-exitChan:
			break Loop
		default:
		}
	}
}
func main() {
	wg.Add(1)
	go f1()
	time.Sleep(5 * time.Second)
	//通知子goroutine退出
	exitChan <- true
	wg.Wait()
}
