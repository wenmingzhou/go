package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周林",
		Age:  9000,
	}
	//序列化
	b, _ := json.Marshal(p1)
	fmt.Println(string(b))

	//反序列化
	str := `{"name":"理想","age":18}`
	var p2 person                    //如果是p2是值拷贝
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在json.Unmarshal内部修改P2的值

	fmt.Println(p2)

	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)

	a1 := 100
	b1 := &a1
	fmt.Println(a1, b1)
	*(b1) = 20
	fmt.Println(a1)
}
