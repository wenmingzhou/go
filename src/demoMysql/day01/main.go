package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//Go连接mysql示例
func main() {
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/goday10"
	//连接数据库
	db, err := sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {
		fmt.Println("dsn格式不正确,", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("open  failed,", err)
		return
	}

}
