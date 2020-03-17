package main

import "fmt"

//要求
//f1(f2)

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}
func main() {
	ret := f3(f2, 100, 200)
	f1(ret)
	var i, j, k int

	fmt.Scanln(&i, &j, &k)
	fmt.Println(i, j, k)
}
