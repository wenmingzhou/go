package main

import "fmt"

var (
	total = 50
	users = []string{"zhou", "liu", "fang", "guo", "fei", "Liu"}
)

//遇到z或者Z分一个金币
//遇到l或者L分二个金币
//遇到g或者f分三个金币
func main() {
	map1 := fenjinbi()
	fmt.Println(map1)
	sum := 0
	for _, val := range map1 {
		sum += val
	}
	fmt.Println(total - sum)
}

func fenjinbi() (map1 map[string]int) {
	map1 = make(map[string]int)
	for i := 0; i < len(users); i++ {
		for _, value := range users[i] {
			c := fmt.Sprintf("%c", value)

			if c == "z" || c == "Z" {
				map1[users[i]] = map1[users[i]] + 1
			} else if c == "l" || c == "L" {
				map1[users[i]] = map1[users[i]] + 2
			} else if c == "g" || c == "f" {
				map1[users[i]] = map1[users[i]] + 3
			}

		}
	}
	return map1
	//fmt.Println(map1)

}
