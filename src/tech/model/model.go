package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/news")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	DB = db
	fmt.Println(DB)

}
