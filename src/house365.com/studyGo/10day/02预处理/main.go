package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //是一个连接池对象

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	//数据库信息 用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:root@tcp(127.0.0.1:3306)/goday10"
	//连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping() //尝试连接数据库
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10) //设置与数据库建立连接的最大数目
	db.SetMaxIdleConns(3)  //设置连接池中的最大闲置连接数
	return
}

//预处理方式插入多条数据
func prepareInsert() {
	//1 写sql语句
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed ,err :%v\n", err)
		return
	}
	defer stmt.Close()
	//后续要拿到stmt去操作
	var m = map[string]int{
		"费静静": 32,
		"李建芬": 32,
		"李建功": 31,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed ,err :%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	prepareInsert()
}
