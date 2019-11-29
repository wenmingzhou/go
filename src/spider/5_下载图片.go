package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var (
	/**
	<img src="http://pic1.win4000.com/pic/c/26/8df58fd858_250_350.jpg" data-original="http://pic1.win4000.com/pic/c/26/8df58fd858_250_350.jpg" alt="巨乳美女睡衣诱惑性感私房写真图片" title="巨乳美女睡衣诱惑性感私房写真图片" style="display: inline;">
	<img class="col-img" src="http://www.veguan.com/uploads/image/20181229/145029f44a9f60cc08b9d3934647a76f.jpg">
	 */
	reImg =`<img[\s\S]+?src="(http[\s\S]+?)"`
)

func getImg(imgurls string) []byte {
	url :=imgurls
	resp,_ :=http.Get(url)
	bytes,_:= ioutil.ReadAll(resp.Body)
	return bytes
}
func main() {
	url :="http://www.veguan.com/ny/"
	resp,_ :=http.Get(url)
	defer resp.Body.Close()
	bytes,_:= ioutil.ReadAll(resp.Body)

	html :=string(bytes)
	//fmt.Println(html)
	re :=regexp.MustCompile(reImg)
	allstring :=re.FindAllStringSubmatch(html,-1)

	num :=len(allstring)
	fmt.Println("图片一个有",num)
	for _,x :=range allstring{
		imgurls :=x[1]
		imageBytes :=getImg(imgurls)
		filename :=`D:\go\pic`+strconv.Itoa(int(time.Now().UnixNano()))+".jpg"
		err :=ioutil.WriteFile( filename,imageBytes,0644)
		if err==nil{
			fmt.Println("下载成功")
		}else{
			fmt.Println("下载失败")
		}
	}
}
