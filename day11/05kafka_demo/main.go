package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 基于sarama 第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()
	// tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要
	// leader和follow都确认
	config.Producer.Partitioner = sarama.NewManualPartitioner //新选出一个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在
	// success channel返回
	// 伪造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log ")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("producer closed,err:%v\n", err)
		return
	}
	fmt.Println("连接成功")
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("sen msg failed,err:%v\n", err)
		return
	}
	fmt.Printf("pid :%v offset:%v\n", pid, offset)

}
