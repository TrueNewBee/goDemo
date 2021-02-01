package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

// tailf的用法示例
// 监听本地文件

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个地方读
		MustExist: false,                                // 文件不存在报错
		Poll:      true,                                 //
	}
	tailes, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail file failed,err:%v\n", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tailes.Lines
		if !ok {
			fmt.Printf("tail file close reopen , filename:%s\n", tailes.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
