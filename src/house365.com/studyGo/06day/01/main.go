package main

import (
	"fmt"
	"os"
)

func f1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed ,err :%v", err)
	}

	fmt.Println(fileObj)
}
func main() {
	f1()
}
