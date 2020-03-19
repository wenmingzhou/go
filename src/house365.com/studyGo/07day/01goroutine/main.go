package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("hello", i)
}
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
