package main

import "fmt"

//结构体是值类型
type person struct {
	name, gender string
}

func f(x *person) {
	(*x).gender = "女" //修改的时副本的gender
	x.gender = "女"    //修改的时副本的gender
}

type Dog struct {
}

func main() {
	var p3 = &person{ //&地址符号，得到结构体指针
		name:   "周文明",
		gender: "男",
	}
	(*p3).name = "刘"
	fmt.Printf("%v\n", p3)
}
