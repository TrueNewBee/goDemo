package main

import (
	"chentianxiang.vip/studygo/logagent/conf"
	"chentianxiang.vip/studygo/logagent/etcd"
	"chentianxiang.vip/studygo/logagent/kafka"
	"chentianxiang.vip/studygo/logagent/tailLog"
	"chentianxiang.vip/studygo/logagent/utils"
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

/*
 使用etcd 的优化版本
*/
// logAgent 入口程序

var (
	cfg = new(conf.AppConf) // 这样声明才有内存地址的指针
)

//func run (){
//	// 1. 读取日志
//	for {
//		select {
//		case line := <-tailLog.ReadLog():
//			// 2. 发送到kafka
//			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
//
//		}
//	}
//
//
//}

func main() {
	// 加载组件
	// 0. 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini faild err:%v\n", err)
	}
	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.ChanMaxSize)
	if err != nil {
		fmt.Printf("init Kafka failed, err %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init Etcd init failed, err %v\n", err)
		return
	}
	fmt.Println("init etcd success")
	// 为了实现每个logagent都拉去自己都有的配置,所以要以自己的IP地址作为区分
	ipStr , err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)

	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("etcd.GetConf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index :%v value:%v\n", index, value)
	}
	//3. 收集日志发往kafka
	tailLog.Init(logEntryConf) // 初始化
	// 因为NewConfChan访问了tskmgr的newConfChan, 这个channel是	tailLog.Init(logEntryConf) 执行初始化的
	//3.1 派一个哨兵去监视日志收集项的变化(有变化及时通知我的logAgent实现热加载配置)
	NewConfChan := tailLog.NewConfChan() // 从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, NewConfChan) // 哨兵发现最新的配置信息会通知上面的哪个通道
	wg.Wait()
	// 4. 具体业务

	//// 2. 打开日志文件准备收集日志
	//err = tailLog.Init(cfg.TaillogConf.FileName)
	//if err != nil{
	//	fmt.Printf("Init taillog failed , err %v\n", err)
	//	return
	//}
	//fmt.Println("init taillog success!")
	//run()
}
