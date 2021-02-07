package tailLog

import (
	"chentianxiang.vip/studygo/logagent/etcd"
	"fmt"
	"time"
)

// laillg 对象的管理者

var tskMgr *tailLogMgr

// tailTask 管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	taskMap     map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集配置信息保存起来
		taskMap:     make(map[string]*TailTask, 32),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
	}
	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		// logEntry.Path : 要收集的日志文件的路径
		// 初始化的时候起了多少个tailtask 都要记下来, 为了后续的判断方便
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.taskMap[mk] = tailObj
	}
	// 后台开启监视
	go tskMgr.run()
}

// 监听自己的newConfChan , 有了新的配置过来之后就做相应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[mk]
				if ok {
					// 原来就有,不需要操作
					continue
				} else {
					// 新增的
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[mk] = tailObj
				}
			}
			// 找到原来 t.logEntry有,但是newConf中没有的, 要删除
			for _, c1 := range t.logEntry { // 从原来的配置中一次拿出配置项
				isDelete := true // 标识位
				for _, c2 := range newConf{ // 从新的配置项中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic{
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					// t.taskMap[mk] == >tailObj
					t.taskMap[mk].cancelFunc()
				}
			}
			fmt.Println("新的配置来了", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数,向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
