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

//查询单个记录
func queryOne(id int) {
	var u1 user
	//查询单条记录的sql语句
	/*
		sqlStr :=`select id ,name ,age from user where id=?;`
		for i:=0;i<12;i++{
			db.QueryRow(sqlStr,1) //从连接池例拿一个
			fmt.Printf("开始第%d次查询\n",i)
		}
	*/
	sqlStr := `select id ,name ,age from user where id=?;`
	rowobj := db.QueryRow(sqlStr, id) //从连接池例拿一个连接取数据库查询
	rowobj.Scan(&u1.id, &u1.name, &u1.age)

	fmt.Println(u1)
}

func queryAll(n int) {

	sqlStr := `select id ,name,age from user where id >?`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("query failed,", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u1 user
		rows.Scan(&u1.id, &u1.name, &u1.age)
		fmt.Println(u1)
	}

}

func insertRow() {
	sqlStr := `insert into user (name,age) values ("周文明",28)`
	res, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("insert failed,", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("get id failed,", err)
		return
	}
	fmt.Println(id)
}

func updateRow(newAge int, id int) {
	sqlStr := `update user set age =? where id =?`
	res, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Println("update failed,", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed,", err)
		return
	}
	fmt.Println(n)
}

func deleteRow(id int) {
	sqlStr := `delete from user  where id =?`
	res, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed,", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed,", err)
		return
	}
	fmt.Println(n)
}

//查询
type user struct {
	id   int
	name string
	age  int
}

//预处理方式插入多条数据
func prepareInsert() {
	//把SQL语句先发给mysql先编译一下,后面再传数据
	sqlStr := `insert into user (name,age) values (?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed,", err)
		return
	}
	defer stmt.Close()
	var m = map[string]int{
		"周嘉琪": 8,
		"周嘉妮": 1,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

//Go连接mysql示例
func main() {
	fmt.Println("1111")
	err := initDb()
	if err != nil {
		fmt.Println("init  db failed,", err)
		return
	}
	fmt.Println("连接数据库成功!")
	//queryOne(2)
	//queryAll(0)
	//insertRow()
	//updateRow(100,3)
	//deleteRow(3)
	prepareInsert()
}
