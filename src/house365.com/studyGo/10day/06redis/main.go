package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

//redis
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("connect redis failed,err :%v\n", err)
		return
	}
	fmt.Println("连接redis成功")

	//#################set/get示例###################
	err = rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)
}
