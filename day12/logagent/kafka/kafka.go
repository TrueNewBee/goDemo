package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// 专门往kafka写日志的模块
// "127.0.0.1:9092"
type logData struct {
	topic string
	data string
}

var (
	client sarama.SyncProducer // 声明一个全局的kafka的生产者client
	logDataChan chan *logData	// 声明一个通道的全局
	)


//Init 初始化client
func Init(addrs []string, maxSize int) (err error) {
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
	// 初始化lodDataChan
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine从通道中取数据发往kafka
	go sendToKafka()
	return
}

// 给外部暴露的一个函数,该函数只有把日志数据发送到一个内部的channel中
func SendToChan(topic, data string){
	msg := &logData{
		topic: topic,
		data:data,
	}
	logDataChan <- msg
}

// 真正向kafka发送消息
func sendToKafka()  {
	for{
		select {
			case ld := <- logDataChan:
				// 构造一个消息
				msg := &sarama.ProducerMessage{}
				msg.Topic = ld.topic
				msg.Value = sarama.StringEncoder(ld.data)
				// 发送到kafka
				pid, offest ,err := client.SendMessage(msg)
				if err != nil {
					fmt.Println("send msg failed, err: ", err)
					return
				}
				fmt.Printf("pid :%v offset:%v\n", pid, offest)
		default:
			time.Sleep(time.Millisecond*50)
		}
	}
}



