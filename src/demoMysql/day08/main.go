package main

import (
	"fmt"
	"time"
)

var exitChan = make(chan bool, 1)

func f() {
LOOP:
	for {
		fmt.Println("周林")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
}
func main() {
	go f()
	time.Sleep(time.Second * 5)
	//如果通知子goroutine退出
	exitChan <- true

}
