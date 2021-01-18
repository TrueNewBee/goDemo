# day06课上比较

# 内容回顾

## 包

包的定义 --> `package`关键字,包名通常是和目录名一致,不能包含`-`

- 一个文件夹就是一个包
- 文件夹中放的都是`.go`文件

包的导入---> `import`

- 包导入路径是从`$GOPATH/src`后面的路径开始写起
- 单行导入
- 多行导入
- 给导入的包起别名
- 匿名导入--->`sql`包导入时候会讲
- Go不支持循环导入

包中`标识符`(变量名\函数名\结构体名\接口名\常量...)可见性 ----> 标识符首字母大写表示对外可见

`init()`

- 包导入的时候会自动执行
- 一个包里面只有一个init()
- init()没有参数也没有用返回值也不能调用它
- 多个包的`init`执行顺序
- 一般用于做一些初始化操作....



## 接口

### 类型断言

![image-20210117131550231](D:\Go\src\chentianxiang.vip\studygo\day06\README.assets\image-20210117131550231.png)

## [文件操作](https://www.liwenzhou.com/posts/Go/go_file/)

### 打开文件和关闭文件 

**关闭文件写的顺序**

![image-20210117132530572](D:\Go\src\chentianxiang.vip\studygo\day06\README.assets\image-20210117132530572.png)

### 读文件

fileObj.Read()

bufio

ioutil

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件
func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	// var temp = make([]byte, 128)	// 指定读的长度
	var temp [128]byte
	for {
		n, err := fileObj.Read(temp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Printf(string(temp[:n]))
		if n < 128 {
			return
		}
	}
}

// 利用bufio这个包按行读文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 创建一个从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed , err:%v", err)
			return
		}
		fmt.Print(line)
	}
}
// 直接读取整个文件件
func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFile1()
	// readFromFilebyBufio()
	readFromFileByIoutil()
}

```



### 写文件

os.OpenFile()

fileObj.write/ filewObj.writeString

bufio

ioutil

```go
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

```

### 小插曲  如何在某个文件中某一行插入内容

1. os.seek() 用该方法,把光标移动到要插入位置
2. 把光标前内容写入到临时文件中
3. 再把想插入内容追加到临时文件中
4. 把光标后面内容追加到临时文件中
5. os.Rename() 把临时文件改成源文件名就好了

# 今日内容

## time 标准库

![image-20210117152437092](D:\Go\src\chentianxiang.vip\studygo\day06\README.assets\image-20210117152437092.png)

## 日志库

  ### 需求分析

   1.支持往不同的地方输出日志

2. 日志级别:

   1. Debug
   2. Trace
   3. Info
   4. Warning
   5. Error
   6. Fatal

3. 日志要支持开关控制, 比如说开发的什么级别都能输出,但上线之后只有INFO级别往下的才能输出 

4. 完整的日志记录要包含 时间,行号,文件名,日志级别,日志信息

5. 日志文件要切割

反射