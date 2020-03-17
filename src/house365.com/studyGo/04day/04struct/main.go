package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string   //性别
	hobby  []string //爱好

}

func main() {
	var p person
	p.name = "周文明"
	p.age = 32
	p.gender = "男"
	p.hobby = []string{"足球", "电影"}
	fmt.Println(p)
}
