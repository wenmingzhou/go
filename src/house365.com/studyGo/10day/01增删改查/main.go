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

func queryOne(id int) {
	var u1 user
	//1 写查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id=?`
	//2 执行
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age) //从连接池里拿出一个连接出来去数据库查询单条记录
	//3 拿到结果
	//rowObj.Scan(&u1.id,&u1.name,&u1.age)  //必须对rowObj对象调用scan方法,会释放数据库连接
	fmt.Printf("u1:%#v", u1)
}

func queryMore(n int) {
	//1 sql语句
	sqlStr := `select id,name,age from user where id>?`
	//2 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed ,err :%v\n", sqlStr, err)
		return
	}
	//3 一定要关闭rows
	defer rows.Close()
	//4 循环取值
	u := make([]user, 0)
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed ,err :%v\n", err)
			return
		}
		u = append(u, u1)
	}
	fmt.Printf("ts:%#v\n", u)
}

func insert() {
	//1 写sql语句
	sqlStr := `insert into user(name,age) values("涂朝阳",28)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed ,err :%v\n", err)
		return
	}
	lastId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("last id ,err :%v\n", err)
		return
	}
	fmt.Printf("id is %v\n", lastId)

}

//更新操作
func updateRow() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 32, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func deleteRow() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 8)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed ,err :%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	queryMore(0)
	insert()
	updateRow()
	deleteRow()
}
