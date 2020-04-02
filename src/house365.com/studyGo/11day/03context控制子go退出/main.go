package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f2(ctx context.Context) {
	defer wg.Done()
Loop:
	for {
		fmt.Println("刘静")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done(): //当读到空的结构体就跳出
			break Loop
		default:
		}
	}
}

func f1(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
Loop:
	for {
		fmt.Println("周文明")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done(): //当读到空的结构体就跳出
			break Loop
		default:
		}
	}
}

//专门用来简化 对于处理单个请求的多个 goroutine 关闭操作
//主goroutine 优雅的通知子goroutine 结束
func main() {
	wg.Add(1)
	// WithCancel 接受一个 Context 并返回一个 Context，返回的 ctx 里面的 Done()
	// 函数会返回一个 chan，当 cancel 调用时，这个 chan 会被关闭。还有一个特性，
	// 当 parent 的 Done() 关闭的时候，孩子 ctx 的 Done() 也会被关闭。
	ctx, cancel := context.WithCancel(context.Background()) //父节
	go f1(ctx)
	time.Sleep(5 * time.Second)
	cancel() //当cancel调用时，写了一个空的结构体  通知子goroutine退出
	wg.Wait()
}
