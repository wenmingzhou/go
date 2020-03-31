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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed ,err :%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")

	var u user
	sqlStr1 := `select id,name,age from user where id=1`
	db.Get(&u, sqlStr1) //查询单条

	fmt.Println(u)

	var userList = make([]user, 0, 10) //必须初始化
	sqlStr2 := `select id,name,age from user`
	db.Select(&userList, sqlStr2)
	fmt.Println(userList) //查询多条
}
