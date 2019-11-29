package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	/**
	<a class="title fl" href="http://hf.sell.house365.com/s_6191129.html" target="_blank">标准两房 房型正 景观房 市政公园旁 商圈繁华 升价空间大</a>
	 */
	reUrl =`<a[\s\S]+?href="(http[\s\S]+?)"`
)
func main() {
	url :="http://nj.sell.house365.com/district/"
	resp,_ :=http.Get(url)
	bytes,_:= ioutil.ReadAll(resp.Body)
	html :=string(bytes)
	re :=regexp.MustCompile(reUrl)
	allstring :=re.FindAllStringSubmatch(html,-1)


	for _,x :=range allstring{
		fmt.Println(x[1])
	}
}
