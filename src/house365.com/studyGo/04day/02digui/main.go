package main

import "fmt"

func f(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
func main() {
	ret := f(5)
	fmt.Println(ret)
}
