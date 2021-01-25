package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http Client
// 在go中全局变量就是一个单例
// 定义一个全局的cline 每次请求就用这一个

// 公用一个clinet适用于 请求比较频繁
var (
	cline = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9090/01")
	if err != nil {
		fmt.Printf("get url failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close() // 一定要记得关闭
	// 从resp中把服务端返回的数据读出
	b, err := ioutil.ReadAll(resp.Body) // 我在客户端读出服务端返回的相应body
	if err != nil {
		fmt.Printf("read resp.Body failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
