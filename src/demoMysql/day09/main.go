package main

import (
	"context"
	"fmt"
	"time"
)

var exitChan = make(chan bool, 1)

func f2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("保德路")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func f(ctx context.Context) {
	go f2(ctx)
LOOP:
	for {
		fmt.Println("周林")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx)
	time.Sleep(time.Second * 5)
	//如果通知子goroutine退出
	cancel()
}
