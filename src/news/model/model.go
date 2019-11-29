package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)
var Db *sqlx.DB
func init() {
	db,err :=sqlx.Open(`mysql`,`root:root@tcp(127.0.0.1:3306)/news?charset=utf8&parseTime=true`)
	//fmt.Println(db,err)
	if err !=nil{
		log.Fatalf(err.Error())
	}
	Db =db
}
