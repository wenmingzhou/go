package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB //是一个连接池对象

type user struct {
	Id   int
	Name string
	Age  int
}

func initDB() (err error) {
	//数据库信息 用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:root@tcp(127.0.0.1:3306)/goday10"
	//连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10) //设置与数据库建立连接的最大数目
	db.SetMaxIdleConns(3)  //设置连接池中的最大闲置连接数
	return
}

// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed ,err :%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")

	sqlInjectDemo("xxx' or 1=1#")                    //SQL:select id, name, age from user where name='xxx' or 1=1#' #号后面当成注释了
	sqlInjectDemo("xxx' union select * from user #") //SQL:select id, name, age from user where name='xxx' union select * from user #'
}
