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

// 插入数据
func insert() {
	// 1. 写SQl语句
	sqlStr := `insert into user(name, age) values("涂朝阳",26)`
	// 2. exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	// 如果是插入的数据,能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id :", id)
}

// 查询多条记录
func queryMore(n int) {
	// 1. SQL语句
	sqlStr := `select id, name, age from user where id >?;` //查询id大于?的
	// 2. 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s,err:%v\n", sqlStr, err)
		return
	}
	// 3, 一定要关闭rows
	defer rows.Close()
	// 4. 循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.name, &u1.id, &u1.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

// 更新操作
func updateRow(newAge int, id int) {
	sqlStrl := `update user set age= ? where id= ?`
	ret, err := db.Exec(sqlStrl, newAge, id)
	if err != nil {
		fmt.Printf("updata,err:%v\n", err)
		return
	}
	line, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据", line)
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

// 删除操作
func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

// 预处理方式 插入多条数据  (批量插入)
func prepareInsert() {
	sqlStr := `insert into user(name, age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 后续只需拿到stmt去执行一些操作
	var m = map[string]int{
		"刘启强": 30,
		"王祥吉": 31,
		"田硕":  32,
		"白慧杰": 55,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	// queryOne(1)
	// queryMore(0)
	// insert()
	// updateRow(8999, 5)
	// deleteRow(2)
	prepareInsert()
}
