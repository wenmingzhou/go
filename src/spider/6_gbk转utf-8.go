package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
import "github.com/axgle/mahonia"

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
func main() {

	url :="https://www.163.com/"
	resp,_ :=http.Get(url)
	bytes,_:= ioutil.ReadAll(resp.Body)
	html :=string(bytes)
	str := html
	str = ConvertToString(str, "gbk", "utf-8")
	fmt.Println(str)

}