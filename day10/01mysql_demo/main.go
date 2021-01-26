package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
)

// Go连接MySQL示例

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/goday10"
	// 连接数据库
	db, err := sql.Open("mysql", dsn) // 不会校验用户和密码正确性
	if err != nil {                   // dsn 格式不正确会报错
		fmt.Printf("dsn %s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	fmt.Println("数据库连接成功")
}
