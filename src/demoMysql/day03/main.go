package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //是一个连接池对象

func initDb() (err error) {
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/goday10"
	//连接数据库
	//db 全局的db
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10)   //设置数据库连接池的最大连接数
	db.SetConnMaxLifetime(5) //设置最大空闲连接数
	return
}

func transactionDemo2() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed,err ", err)
		return
	}
	//执行多个操作
	sqlstr1 := `update user set age=age-2 where id=4`
	_, err = tx.Exec(sqlstr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql3出错了,要回滚啦! ")
		return
	}
	sqlstr2 := `update xxx set age=age+2 where id=5`
	_, err = tx.Exec(sqlstr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql4出错了,要回滚啦! ")
		return
	}

	//上面两步都执行成功,就提交本次事务
	tx.Commit()

	fmt.Println("事务执行成功")
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed,err ", err)
		return
	}
	//执行多个操作
	sqlstr1 := `update user set age=age-2 where id=1`
	_, err = tx.Exec(sqlstr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql1出错了,要回滚啦! ")
		return
	}
	sqlstr2 := `update user set age=age+2 where id=2`
	_, err = tx.Exec(sqlstr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql2出错了,要回滚啦! ")
		return
	}

	transactionDemo2()
	//上面两步都执行成功,就提交本次事务
	tx.Commit()

	fmt.Println("事务执行成功")
}

//查询
type user struct {
	id   int
	name string
	age  int
}

//Go连接mysql示例
func main() {
	err := initDb()
	if err != nil {
		fmt.Println("init  db failed,", err)
		return
	}
	fmt.Println("连接数据库成功!")
	transactionDemo()

}
