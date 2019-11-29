package main

import "fmt"

func main() {
	str :="AFADFDAFAREAFDAFCZVRFSWTGDAFIHDKADFJLAFJDLA"
	map1 :=make(map[string]int)
	for i:=0;i<len(str);i++{
		val,ok :=map1[string(str[i])]
		if ok{
			val++
		}else{
			val =1
		}
		map1[string(str[i])] =val
	}
	fmt.Println(map1)
}
