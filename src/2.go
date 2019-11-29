package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Class struct {
	Id int64
	Name string
	Desc string
}
func main() {
	db,err :=sqlx.Open(`mysql`,`root:root@tcp(127.0.0.1:3306)/news?charset=utf8&parseTime=true`)
	//fmt.Println(db,err)
	mod:=Class{}
	log.Println("----",mod)
	db.Get(&mod,"select * from class")
	log.Println("----",mod)
}
