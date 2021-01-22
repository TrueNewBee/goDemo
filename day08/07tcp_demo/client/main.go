package main

import (
	"fmt"
	"net"
	"os"
)

// tcp client

func main() {
	//1. 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	defer conn.Close()
	if err != nil {
		fmt.Printf("dial 127.0.0.1:20000 failed,err:%v\n", err)
		return
	}
	// 2. 发送数据
	var msg string
	if len(os.Args) < 2 {
		msg = "hello wang ye "
	} else {
		msg = os.Args[1]
	}
	conn.Write([]byte(msg))
}
