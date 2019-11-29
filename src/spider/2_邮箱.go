package main

import (
	"fmt"
	"regexp"
)

var (
	reEail =`[\w\.]+@\w+\.[a-z]{2,3}(\.[a-z]{2,3})?`
)
func main() {

	html :="524797132@qq.com;"
	html +="52472@qq.com;"
	html +="529713211@qq.com;"
	html +="52479345@qq.com;"
	html +="52479345@163.com;"
	html +="zhouwm@163.com;"
	html +="zhouwm@163.com.cn;"
	html +="zhou.wm@163.com.cn;"
	re :=regexp.MustCompile(reEail)
	allstring :=re.FindAllStringSubmatch(html,-1)


	for _,x :=range allstring{
		fmt.Println(x[0])
	}
}
