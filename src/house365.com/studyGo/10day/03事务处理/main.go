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

func transactionDemo() {
	//1 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed ,err :%v\n", err)
		return
	}
	//执行多个sql操作
	sqlStr1 := `update user set age=age-2 where id=1`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Printf("执行sql1出错啦，要回滚 !")
		return
	}
	sqlStr2 := `update user set age=age+2 where id1=3` // 当id1不存在时 要回滚
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Printf("执行sql2出错啦，要回滚 !")
		return
	}
	//上面两步sql都执行成功,就提交本次事务
	err = tx.Commit()
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Printf("提交出错了，要回滚!")
		return
	}
	fmt.Printf("事务执行成功!")
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed ,err :%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")

	transactionDemo()

}
