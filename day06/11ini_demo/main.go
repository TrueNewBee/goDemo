package main

import "fmt"

// ini配置文件解析器

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"prot"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func loadIni(v interface{}) {
	// 先打开文件
	// 通过反射查找字段
	// 给相应字段赋值.
	// 关于反射可以看前几个reflect_demo1/reflect_demo2
}

func main() {
	var mc MysqlConfig
	loadIni(&mc)
	fmt.Println(mc.Address, mc.Password, mc.Username, mc.Port)
}
