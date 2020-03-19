package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
  使用goroutine和channel实现一个计算int64随机数各位数和的程序。
	1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	3.主goroutine从resultChan取出结果并打印到终端输出
*/

type result struct {
	value int64
	sum   int64
}

var jobChan = make(chan int64, 100)
var resultChan = make(chan result, 100)
var wg sync.WaitGroup

func send() {
	defer wg.Done()
	for {
		x := rand.Int63()
		jobChan <- x
		time.Sleep(time.Second)
	}
}

func get() {
	defer wg.Done()
	for {
		val := <-jobChan
		n := val
		sum := int64(0)
		for n > 0 {
			sum += n % 10
			n = n / 10
		}

		newResult := result{
			value: val,
			sum:   sum,
		}

		resultChan <- newResult

	}
}
func main() {
	wg.Add(1)
	go send()
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go get()
	}
	for value := range resultChan {
		fmt.Println(value)
	}
	wg.Wait()
	//jobChan :=make(chan int64)
}
