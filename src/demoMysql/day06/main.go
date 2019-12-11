package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

//全局的redis对象(客户端)
var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}
func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("init redis failed,", err)
		return
	}
	fmt.Println("连接Redis成功")

	//set/get示例

	err = redisdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Println("set redis failed,", err)
		return
	}

	val, _ := redisdb.Get("score").Result()
	fmt.Println(val)

	//zset 示例

	key := "rank"
	items := []*redis.Z{ //结构体指针
		&redis.Z{90, "php"},
		&redis.Z{96, "go"},
		&redis.Z{97, "python"},
		&redis.Z{99, "java"},
	}
	//把原始都追加到key
	redisdb.ZAdd(key, items...)

	//给制定原始加分数
	newScore, err := redisdb.ZIncrBy(key, 10.0, "go").Result()
	fmt.Println(newScore)
}
