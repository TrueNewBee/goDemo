package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
)

// Go连接MySQL示例

var db *sql.DB // 是一个连接池对象  单例 全局的

type user struct {
	id   string
	name string
	age  int
}

// 初始化数据库连接
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

// 查询单条记录
func queryOne(id int) {
	var u1 user
	// 1.写入查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id=?`
	// 执行直接操作
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	// 打印结果
	fmt.Printf("u1:%#v\n", u1)
	// 执行过程
	// rowObj := db.QueryRow(sqlStr, 1) //	从连接池拿一个连接出来去数据库查询单条记录
	// rowObj.Scan(&u1.id, &u1.name, &u1.age) // 必须对rowObj对象调用Scan方法,因为该方法会释放数据库连接
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	queryOne(1)
}
