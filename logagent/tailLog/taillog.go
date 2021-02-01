package tailLog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

// 专门从日志文件收集日志的模块

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个地方读
		MustExist: false,                                // 文件不存在报错
		Poll:      true,                                 //
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail file failed,err:%v\n", err)
		return
	}
	return
}

func ReadLog() <-chan *tail.Line{
	return tailObj.Lines
}