package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
	"github.com/jmoiron/sqlx"
)

// Go连接MySQL示例

var db *sqlx.DB // 是一个连接池对象  单例 全局的

// sqlx用的反射,故需要大写
type user struct {
	ID   int
	Name string
	Age  int
}

// 初始化数据库连接
func initDB() (err error) {
	// 连接数据库 用户名:密码@@tcp(ip:端口号)/数据库名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库 (去掉:= 声明  因为已经var 过了)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 最大空闲连接数
	db.SetMaxIdleConns(2)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	sqlStr1 := `select id, name, age from user where id=1`
	var u user
	// 单条查询
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	// 多条查询
	var userList = make([]user, 0, 10)
	sqlStr2 := `select id , name , age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)
}
