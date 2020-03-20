package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func send() {
	defer wg.Done()
	ch := make(chan int, 1)
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(time.Second * 2)
	}
}

func main() {

	wg.Add(1)
	go send()
	fmt.Println("over")
	wg.Wait()

}
