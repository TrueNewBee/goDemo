package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容
// 0100 0000

func writeDemo1() {
	// 一些或 | 的参数
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()
	// write
	fileObj.Write([]byte("zhoulin mengbi le!"))
	// writeString
	fileObj.WriteString("周林解释不了了!\n")
}

func writeDemo2() {
	// 一些或 | 的参数
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello 沙河\n") // 写到缓存中
	wr.Flush()                   // 将缓存中的内容写入到文件
}
func writeDemo3() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file faild , err:", err)
		return
	}
}

func main() {
	// writeDemo1()
	// writeDemo2()
	writeDemo3()
}
