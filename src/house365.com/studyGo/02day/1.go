package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
func main() {
	a := 10
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)

	var p *int
	fmt.Println(p)

	var p2 = new(int)
	fmt.Println(p2)
	*p2 = 200
	fmt.Println(p2)
	fmt.Println(*p2)

	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["理想"] = 18
	fmt.Println(m1)

	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	var keys = make([]string, 0, 200)

	for key := range scoreMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	var s1 = make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][80] = "周"
	s1[0][70] = "刘"
	fmt.Println(s1)

	var m2 = make(map[string][]int, 10)
	m2["北京"] = []int{10, 20, 30}
	fmt.Println(m2)

	str1 := "how do you do"
	arr1 := strings.Split(str1, " ")
	fmt.Println(arr1)

	map2 := make(map[string]int, 10)
	for _, v := range arr1 {
		if _, ok := map2[v]; !ok {
			map2[v] = 1
		} else {
			map2[v]++
		}
	}
	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Println(key, value)
	}

	//回文判断
	ss := "上海自来水来自海上" //把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune, 0)
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println(r)
	//for i,_ :=range ss{
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
