package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	str := "hello 沙河"
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求,参数都放在URL上(query param), 请求体重是没有数据的
	queryParam := r.URL.Query() // 自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	fmt.Println(name)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) // 我在服务端打印客户端发来的请求
	w.Write([]byte("ok"))

}

func main() {
	http.HandleFunc("/01", f1)
	http.HandleFunc("/02", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
