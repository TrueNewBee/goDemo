package main

import (
	"fmt"
	"net"
)

// tcp server 端
func processConn(conn net.Conn) {
	defer conn.Close()
	// 3.与客户端通信
	var temp [128]byte
	n, err := conn.Read(temp[:])
	if err != nil {
		fmt.Printf("read from conn failed,err:%v\n", err)
		return
	}
	fmt.Println(string(temp[:n]))
}

func main() {
	// 1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("start tcp server on  127.0.0.1:20000,err:%v\n", err)
		return
	}
	for {
		// 2. 等待别人来跟我建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			return
		}
		go processConn(conn)
	}

}
