package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串中解析出对应的数据
	str := "1000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parseint failed ,err:%v\n", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)
	// Atoi: 把字符串转换成 int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 从字符串中解析出浮点数
	floatStr := "1.23456"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)

	// 把数字转换成字符串类型
	i := 97
	// ret2 := string(i)	// "a"
	ret2 := fmt.Sprintf("%d", i) // "97"
	fmt.Printf("%#v", ret2)
	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v", ret3)
}
