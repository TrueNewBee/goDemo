package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

// 1. 把Go语言中的结构体变量 --> json格式的字符串
// 2. json格式的字符串 ---> Go语言中能够识别的结构体变量

type person struct {
	// 变量名大写,json包才能访问
	// `json:"name"`  把在json中显示小写
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周林",
		Age:  1888,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	fmt.Printf(string(b))

	// 反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	// 传入指针是为了json.Unmarshal内部修改p2
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
	fmt.Printf("%v\n", p2)

}
