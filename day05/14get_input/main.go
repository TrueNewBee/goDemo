package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格

func useScan() {
	var s string
	fmt.Println("请输入内容:")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是:%s\n", s)
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin) // 标准输出
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是: %s\n", s)

}
func main() {
	useBufio()

	// 作业
	// logger.Trace()
	// logger.Debug()
	// logger.Warning()
	// logger.Info()
	// logger.Error("日志的内容")

	// 写日志
	fmt.Fprintln(os.Stdout, "只是一条日志记录") // 标准输出
	// 写到指定文件
	fileObj, _ := os.OpenFile("./xx.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	fmt.Fprintln(fileObj, "这是一条日志记录")
}
