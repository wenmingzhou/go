package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())

	ret := time.Unix(1584337792, 0)
	fmt.Println(ret)

	ret2 := time.Unix(1584337792, 10)
	fmt.Println(ret2)
}
