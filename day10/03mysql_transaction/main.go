package main

import (
	"database/sql"
	"fmt"
	// _ "github.com/go-sql-driver/mysql" //init()
)

var db *sql.DB // 是一个连接池对象  单例 全局的v

type user struct {
	id   string
	name string
	age  int
}

func initDB() (err error) {
	// 连接数据库 用户名:密码@@tcp(ip:端口号)/数据库名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库 (去掉:= 声明  因为已经var 过了)
	db, err = sql.Open("mysql", dsn) // 不会校验用户和密码正确性
	if err != nil {                  // dsn 格式不正确会报错
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 最大空闲连接数
	db.SetMaxIdleConns(2)
	return
}

// MySQL事务
func transactionDemo() {
	//1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin,err:%v\n", err)
		return
	}
	// 执行多个sql操作
	sqlstr1 := `update user set age=age-2 where id=1`
	sqlstr2 := `update user set age=age+2 where id=3`
	// 执行sql1
	_, err = tx.Exec(sqlstr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL1出错了,要回滚!")
		return
	}
	// 执行sql2
	_, err = tx.Exec(sqlstr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL2出错了,要回滚!")
		return
	}
	// 上面两步SQL都执行成功,就提交本次事务
	err = tx.Commit()
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("提交出错了,要回滚!")
		return
	}
	fmt.Println("事务执行通过")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	transactionDemo()
}
