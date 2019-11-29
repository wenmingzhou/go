package main

import (
	"fmt"
	"time"
)

func main() {
	s3 :="1999年10月10日"
	t3,_ :=time.Parse("2006年01月02日",s3)
	fmt.Println(t3)
}
