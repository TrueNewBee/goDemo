package tailLog

import (
	"chentianxiang.vip/studygo/logagent/kafka"
	"context"
	"fmt"
	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了实现能退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// NewTailTask 构造函数
func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init() // 根据路径去打开对应的日志
	return tailObj
}

// 打开相应的日志
func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个地方读
		MustExist: false,                                // 文件不存在报错
		Poll:      true,                                 //
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Printf("tail file failed,err:%v\n", err)
	}

	// 当goroutine执行的函数退出的后, goroutine 就退出了
	go t.run() // 直接去采集日志发送到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done(): // 有值的时候就退出了
			fmt.Printf("tail task:%s 结束了\n", t.path+t.topic)
			return

		case line := <-t.instance.Lines: // 从tailObj的通道中一行一行的读取日志数据
			// 发往 kafka
			//kafka.SendToKafka(t.topic, line.Text)	// 函数调用函数 同步了
			// 先把日志数据发到一个通道中 这样杜绝的函数调用函数比较快
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
