package main

import "fmt"

type person struct {
	name, gender string
	age          int
}

//使用值接受者:传拷贝
func (p person) guonian() {
	p.age++
}

//指针接受者:传内存地址进去
func (p *person) zhenguonian() {
	p.age++
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPerson("周文明", 32)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.age) //====>fmt.Println((*p1).age)
	p1.guonian()        //====>(*p1).guonian()
	fmt.Println(p1.age)

	p1.zhenguonian() //p1的地址传递
	fmt.Println(p1.age)
}
