package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB //是一个连接池对象

func initDb() (err error) {
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/goday10"
	//连接数据库
	//db 全局的db
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10)   //设置数据库连接池的最大连接数
	db.SetConnMaxLifetime(5) //设置最大空闲连接数
	return
}

//查询
type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//Go连接mysql示例
func main() {
	err := initDb()
	if err != nil {
		fmt.Println("init  db failed,", err)
		return
	}
	fmt.Println("连接数据库成功!")

	sqlStr1 := `select id ,name,age from user where id=1`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Println(u)

	var userList = make([]user, 0)
	sqlStr2 := `select id ,name,age from user`
	db.Select(&userList, sqlStr2)
	fmt.Println(userList)
}
