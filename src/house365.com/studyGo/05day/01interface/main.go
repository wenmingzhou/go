package main

import "fmt"

//interface是一组method的集合
//定义一个speaker接口
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}
type cat struct{}

type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (c dog) speak() {
	fmt.Println("汪汪汪")
}

func da(x speaker) {
	x.speak()
}

func main() {
	var c cat
	var d dog
	//c.speak()
	da(c)
	da(d)

	//一个变量如果实现了接口中规定的所有的方法，那么这个变量就实现了这个接口
	//可以称为为这个接口类型的变量   dog,cat 称为speaker类型变量
	var ss speaker //定义一个接口类型 speaker 的变量 ss

	fmt.Printf("%T\n", ss)
	ss = c //因为dog类型实现了speaker所有的方法，所以可以赋值
	fmt.Printf("%T\n", ss)

	ss = d
	fmt.Printf("%T\n", ss)
}
