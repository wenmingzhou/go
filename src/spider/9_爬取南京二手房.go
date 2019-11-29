package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var tickts =264980590
var wg sync.WaitGroup
var mu sync.Mutex
func main() {
	wg.Add(100)
	for i:=1;i<=100;i++ {
		go saleTickts("爬虫:",i)
	}
	wg.Wait()
	fmt.Println("爬取完成...")
}

func saleTickts(name string,i int)  {
	rand.Seed(time.Now().UnixNano())
	for{
		//mu.Lock()
		if tickts>0{
			time.Sleep(time.Duration(rand.Intn(1000)))

			url :="http://nj.sell.house365.com/s_"+strconv.Itoa(tickts)+".html?spiderpage"
			http.Get(url)
			fmt.Println(name,i,"正在爬取页面",url)
			tickts--
		}else{
			fmt.Println(name,",爬取结束了")
			//mu.Unlock()
			break
		}
		//mu.Unlock()
	}
	wg.Done()
}
