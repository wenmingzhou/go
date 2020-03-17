package main

import "fmt"

func main() {
	lixiang(f2, 100, 200)
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
func lixiang(x func(int, int), m, n int) {
	tmp := func() {
		x(m, n)
	}
	tmp()
}
