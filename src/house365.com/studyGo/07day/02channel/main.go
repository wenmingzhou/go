package main

import "fmt"

var a []int

func main() {
	b := make(chan int)
	fmt.Println(b)
	go func() {
		x := <-b
		fmt.Println(x)
	}()
	b <- 10
}
