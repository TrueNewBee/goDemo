package main

import (
	"chentianxiang.vip/studygo/logagent/conf"
	"chentianxiang.vip/studygo/logagent/kafka"
	"chentianxiang.vip/studygo/logagent/tailLog"
	"fmt"
	"gopkg.in/ini.v1"
)
// logAgent 入口程序

var (
	cfg =new(conf.AppConf) // 这样声明才有内存地址的指针
)

func run (){
	// 1. 读取日志
	for {
		select {
		case line := <-tailLog.ReadLog():
			// 2. 发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)

		}
	}


}

func main() {
	// 加载组件
	// 0. 加载配置文件
	//cfg,err := ini.Load("./conf/config.ini")
	//fmt.Println(cfg.Section("kafka").Key("address"))
	//fmt.Println(cfg.Section("kafka").Key("topic"))
	//fmt.Println(cfg.Section("taillog").Key("path"))
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil{
		fmt.Printf("load ini faild err:%v\n", err)
	}
	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil{
		fmt.Printf("init Kafka failed, err %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2. 打开日志文件准备收集日志
	err = tailLog.Init(cfg.TaillogConf.FileName)
	if err != nil{
		fmt.Printf("Init taillog failed , err %v\n", err)
		return
	}
	fmt.Println("init taillog success!")
	run()
}