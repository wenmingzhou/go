package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	rePhone =`1[3456789]\d{9}`
)
func main() {
	url :="https://www.haomagujia.com/"
	resp,_ :=http.Get(url)
	bytes,_:= ioutil.ReadAll(resp.Body)
	html :=string(bytes)
	re :=regexp.MustCompile(rePhone)
	allstring :=re.FindAllStringSubmatch(html,-1)


	for _,x :=range allstring{
		fmt.Println(x[0])
	}
}
