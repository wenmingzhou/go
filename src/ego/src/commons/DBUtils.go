package commons

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	stmt *sql.Stmt
	rows *sql.Rows
	tx   *sql.Tx
)

//事务打开连接
func OpenConnWithTx() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/ego")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	tx, err = db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

//判断返回值是否大于1
func PrepareWithTx(sql string, args ...interface{}) int {
	result, err := tx.Exec(sql, args...)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//提交事务连接
func CloseConnWithTx(result bool) {
	if result {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	if rows != nil {
		rows.Close()
	}
	if stmt != nil {
		stmt.Close()
	}
	if db != nil {
		db.Close()
	}
}

//打开连接
func openConn() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/ego")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	return nil
}

//关闭连接
func CloseConn() {
	if rows != nil {
		rows.Close()
	}
	if stmt != nil {
		stmt.Close()
	}
	if db != nil {
		db.Close()
	}
}

//执行DML新增，删除，修改操作
func Dml(sql string, args ...interface{}) (int64, error) {
	err := openConn()
	if err != nil {
		fmt.Println("执行DML时出现错误,打开数据库失败", err)
		return 0, err
	}
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DML时出现错误,预处理出现错误", err)
		return 0, err
	}
	//此处要有...表示切片,如果没有表示数组，会报错
	result, err := stmt.Exec(args...)
	if err != nil {
		fmt.Println("执行DML时出现错误,执行出现错误", err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("执行DML时出现错误,获取受影响行数错误", err)
		return 0, err
	}
	CloseConn()
	return count, err
}

//执行DQL查询操作
func Dql(sql string, args ...interface{}) (*sql.Rows, error) {
	err := openConn()
	if err != nil {
		fmt.Println("执行DQL出现错误,打开连接失败", err)
		return nil, err
	}
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DQL时出现错误,预处理出现错误", err)
		return nil, err
	}
	//此处要有...表示切片,如果没有表示数组，会报错
	rows, err := stmt.Query(args...)
	if err != nil {
		fmt.Println("执行DQL时出现错误,执行出现错误", err)
		return nil, err
	}

	return rows, nil
}
