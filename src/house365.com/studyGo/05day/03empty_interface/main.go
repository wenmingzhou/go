package main

import "fmt"

//interface 关键字
//ingerface{} 空接口类型
func main() {
	m1 := make(map[string]interface{}, 20)
	fmt.Println(m1)

	assign(100)
}

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str := a.(string)
	fmt.Println(str)
}
