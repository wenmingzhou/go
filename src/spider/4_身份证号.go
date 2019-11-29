package main

import (
	"fmt"
	"regexp"
)

var (
	/**
	321023198711053210
	 */
	//reID =`\d{17}[\d|x]`
	reID =`[1-6]\d{5}((19\d{2})|(20(0\d)|(1[0-9])))\d{7}[\d|x]`
)
func main() {

	html :="321023198711053210;"
	html +="32102319871105321x;"
	html +="431023198711053210;"
	html +="789111111111111111;"
	html +="321023208711053210;"
	re :=regexp.MustCompile(reID)
	allstring :=re.FindAllStringSubmatch(html,-1)


	for _,x :=range allstring{
		fmt.Println(x[0])
	}
}
