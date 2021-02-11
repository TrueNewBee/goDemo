# day13 课上笔记



# 内容回顾

## LogAgent的实现

1. 配置文件版LogAgent实现
   1. kafka的使用
   2. lailf第三方包的使用
   3. ini配置文件解析
2. etcd版本的LogAgent实现
   1. 项目启动时从etcd拉取手机日志项信息
   2. 利用watch实时监控etcd中配置的变化
   3. 利用IP每个LogAgent分别从etcd拉取 自己的配置

## 留两个思考题(自主学习能力)

1. Raft协议
2. watch底层实现原理

# 今日内容

## LogTransfer

从kafka里面把日志取出来,写入ES,使用Kibana做可视化展示

### Elastic search

[ES介绍博客地址](https://www.liwenzhou.com/posts/Go/go_elasticsearch/)

###  使用

下下载(华为云),解压,

![image-20210207164135046](D:\Go\src\chentianxiang.vip\studygo\day13\day13课上笔记.assets\image-20210207164135046.png)

Kibana

![image-20210208104011223](D:\Go\src\chentianxiang.vip\studygo\day13\day13课上笔记.assets\image-20210208104011223.png)

![image-20210208104316155](D:\Go\src\chentianxiang.vip\studygo\day13\day13课上笔记.assets\image-20210208104316155.png)

![image-20210208104428004](D:\Go\src\chentianxiang.vip\studygo\day13\day13课上笔记.assets\image-20210208104428004.png)

### kafka消费

根据topic找所有的分区

每一个分区去消费数据

```go
package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}
```

### LogTransfer实现

### 加载配置文件

```go
// 0.加载配置文件
var cfg = new(conf.LogTransfer) // 这样声明才有内存地址的指针
err := ini.MapTo(cfg, "./conf/cfg.ini")
if err != nil {
	fmt.Printf("init config, err:%v\n", err)
	return
}
fmt.Printf("cfg:%v\n",cfg)
```

两个坑:

1. 在一个函数中修改变量一定要传指针
2. 在配置文件对应的结构体中一定要设置tag(特别是嵌套的结构体)

```go
// LogTransfer 全局配置
type LogTransfer struct {
	KafkaCfg `ini:"kafka"`		// 写上tab
	ESCfg `ini"es"`
}
// KafkaCfg ...
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

// ESCfg ...
type ESCfg struct {
	Address string `ini:"address"`
}
```

 

## 系统监控

gopsutil做系统监控信息的采集,写入influxDB,使用grafana作展示

prometheus监控: 采集性能指标数据,保存起来,使用grafana作展示

![image-20210211183653754](D:\Go\src\chentianxiang.vip\studygo\day13\day13课上笔记.assets\image-20210211183653754.png)



## 项目总结

服务端 Agent开发

1.  项目的架构图
2. 为什么不用ELK
3. logAgent里面如何保证日志不丢/重启之后继续收集日志(记录读取文件的offset)
4. kafka课上整理的那一些
5. etcd的watch的原理
6. es相关知识
7. **项目一定要自己写一遍** 



## 找工作

1. 找工作的话还是算法和数据结构(刷LeetCode)

2. 找运维开发的话前端自己会一点加分,时下热点的技术栈

3. 简历好好写

4. Boss直聘等该花钱花钱

   

## 国庆后

web框架: gin

微服务

Docker和K8s



## 今日分享

工作和生活都很重要.