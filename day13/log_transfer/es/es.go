package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

// 初始化ES, 准备接受kafka那边发来的数据
// LogData ...
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

// Init
func Init(address string,  chanSize  ,nums int) (err error) {
	if strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return err
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogData, chanSize)
	for i :=0;i<nums; i++{
		go senToEs()
	}
	return
}

// SendToES 发送数据到ES
func SendToESChan(msg *LogData) {
	//
	ch <- msg

}
func senToEs() {
	// 链式操作
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("indexed student %s to index %s\n", put1.Id, put1.Index)
		default:
			time.Sleep(time.Second)
		}
	}
}
