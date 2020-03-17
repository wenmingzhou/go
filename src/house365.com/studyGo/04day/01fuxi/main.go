package main

import (
	"fmt"
)

func yuanshui(name string) {
	fmt.Println("hello", name)
}

func low(f func()) {
	f()
}

//yuanshui 函数当成参数传递到low函数中
//闭包
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}
func main() {
	ret := bi(yuanshui, "zhou")
	low(ret)
}
