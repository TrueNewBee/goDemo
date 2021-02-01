package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 专门往kafka写日志的模块
// "127.0.0.1:9092"
var (
	client sarama.SyncProducer // 声明一个全局的kafka的生产者client
)


//Init 初始化client
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	// tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要
	// leader和follow都确认
	config.Producer.Partitioner = sarama.NewManualPartitioner //新选出一个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在
	// success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Printf("producer closed,err:%v\n", err)
		return
	}
	return
}

func SendToKafka(topic, data string)  {
	// 伪造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送到kafka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("sen msg failed,err:%v\n", err)
		return
	}
	fmt.Printf("pid :%v offset:%v\n", pid, offset)
	fmt.Println("发送成功")
}



