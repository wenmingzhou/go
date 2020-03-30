package main

import (
	"fmt"
	"house365.com/studyGo/09day/split_string"
)

func main() {
	ret := split_string.Split("abcbc", "b")
	fmt.Println(ret)

	ret2 := split_string.Split("bbb", "b")
	fmt.Println(ret2)

	ret3 := split_string.Split("mmm", "b")
	fmt.Println(ret3)
}
