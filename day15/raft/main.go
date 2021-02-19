package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

// 1. 实现3节点选举
// 2. 改造代码成分布式选举代码
// 3. 演示完整代码, 自动选举, 日志复制

// 通过Rpc实现了分布式

// 定义3节点常量
const raftCount = 3

// 声明Leader对象
type Leader struct {
	// 任期
	Term int
	// LeaderId编号
	LeaderId int
}

// 声明raft
type Raft struct {
	// 锁
	mu sync.Mutex
	// 节点编号
	me int
	// 当前任期
	currentTerm int
	// 为哪个节点,投票
	votedFor int
	// 3个状态
	// 0 follower 1 candidate 2 leader
	state int
	// 发送最后一条数据的时间
	lastMessageTime int64
	// 设置当前节点的领导
	currentLeader int
	// 节点间发消息的通道
	message chan bool
	// 选举通道
	electCh chan bool
	// 心跳信号的通道
	heartBeat chan bool
	// 返回心跳信号的通道
	heartbeatRe chan bool
	// 超出时间
	timeout int
}

// 0 还没上任 -1 没有编号
var leader = Leader{0, -1}

func main() {
	// 过程: 有3个节点,最初都是follower
	// 若有candidate状态,进行投票拉票
	// 会产生leader

	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		// 创建3个raft节点
		Make(i)
	}
	// 加入服务端监听
	rpc.Register(new(Raft))
	rpc.HandleHTTP()
	// 监听服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
	}
}

func Make(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	// -1 代表谁都不投,此时节点刚创建
	rf.votedFor = -1
	// 0 follower
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	//节点任期
	rf.setTerm(0)
	//// 节点间发消息的通道
	//message chan bool
	//// 选举通道
	//electCh chan bool
	//// 心跳信号的通道
	//heartBeat chan bool
	//// 返回心跳信号的通道
	//hearbeatRe chan bool
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 选举的协程
	go rf.election()

	// 心跳检测的协程
	go rf.sendLeaderHearBeat()
	return rf
}

func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

// 选举
func (rf *Raft) election() {
	// 设置标记,判断是否选出了leader
	var result bool
	for {
		// 设置超时,150到300随机数
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()
		select {
		// 延迟等待1毫秒
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("节点当前状态为:", rf.state)
		}
		result = false
		for !result {
			// 选主逻辑
			result = rf.election_one_round(&leader)
		}
	}
}

// 随机值
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间,发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 实现选主的逻辑
func (rf *Raft) election_one_round(leader *Leader) bool {
	//定义超时
	var timeout int64
	timeout = 100
	// 投票数量
	var vote int
	// 定义是否开始心跳信号的产生
	var triggerHeartbeat bool
	// 时间
	last := millisecond()
	// 用于返回值
	success := false

	// 给当前节点变成cadidate
	rf.mu.Lock()
	//修改状态
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start electing leader")
	for {
		// 遍历所有节点拉选票
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				//拉选票
				go func() {
					if leader.LeaderId < 0 {
						// 设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量
		vote = 0
		// 遍历
		for i := 0; i < raftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-rf.electCh:
				if ok {
					// 投票数量加1
					vote++
					// 若选票个数,大于节点个数/2.则成功
					success = vote > raftCount/2
					if success && !triggerHeartbeat {
						//变化成主节点,选主成功了
						// 开始出发心跳信号检测
						triggerHeartbeat = true
						rf.mu.Lock()
						// 变主
						rf.becomeLeader()
						rf.mu.Unlock()
						// 由leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为了leader")
						fmt.Println("leader开始发送心跳信号了")
					}
				}
			}
		}
		// 做最后的校验工作
		// 若不超时,且票数大于一半,则选举成功, break
		if timeout+last < millisecond() || (vote > raftCount/2 || rf.currentLeader > -1) {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):

			}
		}
	}
	return success
}

// 修改状态 candidate
func (rf *Raft) becomeCandidate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.votedFor = rf.me
	rf.currentLeader = -1 // 选举中没领导
}

// 修改状态 leader
func (rf *Raft) becomeLeader() {
	rf.state = 2
	rf.currentLeader = rf.me
}

// leader 节点发送心跳信号
//  顺便完成数据同步, 这里先不实现,后面有现成的
// 看小弟挂没挂
func (rf *Raft) sendLeaderHearBeat() {
	// 死循环
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}
	}
}

//  用于返回给leader的确认信号
func (rf *Raft) sendAppendEntriesImpl() {
	// 是主就别跑了
	if rf.currentLeader == rf.me {
		// 此时是leader
		// 记录确认信号的节点个数
		var success_cout = 0
		// 设置确认信号
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					//rf.heartbeatRe<-true
					// 这里实际上相当于客户端
					rp, err := rpc.DialHTTP("tcp", "127.0.0.1")
					if err != nil {
						log.Fatal(err)
					}
					// 接受服务器返回的信息
					// 接受服务端返回信息的变量
					var ok = false
					err = rp.Call("Raft.Communication", Param{"hello "}, &ok)
					if err != nil {
						log.Fatal(err)
					}
					if ok {
						rf.heartbeatRe <-true
					}
				}()
			}
		}
		//  计算返回确认信号的个数
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.heartbeatRe:
				if ok {
					success_cout++
					if success_cout > raftCount/2 {
						fmt.Println("投票选举成功,心跳信号ok")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}
}

// 首字母大写,RPC会犯
// 分布式通信
type Param struct {
	Msg string
}

// 通信方法
func (r *Raft) Communication(p Param, a *bool) error {
	fmt.Println(p.Msg)
	*a = true
	return nil
}
