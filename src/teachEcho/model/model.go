package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/news?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	DB = db

}
