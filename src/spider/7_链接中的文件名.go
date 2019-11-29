package main

import (
	"fmt"
	"regexp"
)

var (
	reUrlName =`/(\w+\.((jpg)|(png)))`
)
func main() {

	html :="http://img31.house365.com/M01/7E/59/rBEBqFzeZz2ADsICAABnaVCK5mY062_322x260_c1.jpg;"
	html +="http://img32.house365.com/M01/7E/59/rBEBqFzeZz2ADsICAABnaVCK5mY062_322x260_c2.png;"
	html +="http://img33.house365.com/M01/7E/59/rBEBqFzeZz2ADsICAABnaVCK5mY062_322x260_c3.jpg;"
	html +="http://img34.house365.com/M01/7E/59/rBEBqFzeZz2ADsICAABnaVCK5mY062_322x260_c4.jpg;"
	html +="http://img35.house365.com/M01/7E/59/rBEBqFzeZz2ADsICAABnaVCK5mY062_322x260_c5.jpg;"

	re :=regexp.MustCompile(reUrlName)
	allstring :=re.FindAllStringSubmatch(html,-1)


	for _,x :=range allstring{
		fmt.Println(x[1])
	}
}
