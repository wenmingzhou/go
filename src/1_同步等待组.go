package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
	var wg sync.WaitGroup
	wg.Add(1)
	go pringNum1(&wg)
	go pringNum2(&wg)
	wg.Wait()
	 */
	var wg sync.WaitGroup
	for i:=1;i<3;i++{
		pringNum(&wg,i)
		//fmt.Println(i)
	}
	wg.Wait()
}

func pringNum(wg *sync.WaitGroup,num int)  {

	wg.Add(1)
	for i:=1; i<=100; i++ {
		fmt.Println("子goroutine1:",num,i)
	}
	wg.Done()
}
func pringNum1(wg *sync.WaitGroup)  {
	for i:=1; i<=100; i++ {
		fmt.Println("子goroutine1,i:",i)
	}
	wg.Done()
}
func pringNum2(wg *sync.WaitGroup)  {
	for i:=1; i<=100; i++ {
		fmt.Println("子goroutine2,i:",i)
	}
	wg.Done()
}
