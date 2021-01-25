package main

import (
	"flag"
	"fmt"
	"time"
)

// flag获取命令行参数
// 终端输入 :  05flag.exe  -name 周林  -age 900000 -married=true  -ct=300h
func main() {
	//方法 1 创建一个标志位参数
	age := flag.Int("age", 9000, "请输入名字")
	married := flag.Bool("married", true, "结婚了没")
	cTime := flag.Duration("ct", time.Second, "结婚多久了")

	// 使用flag 解析变量值
	flag.Parse()

	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
	// 方法2
	var name string
	flag.StringVar(&name, "name", "我是默认名字", "请输入名字")

}
