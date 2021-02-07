# day12课上笔记

# 内容回顾

## go module

依赖管理工具

## context

goroutine管理.

`context.Context`

两个根节点. `context.Background()`, `context.TODO()`

四个方法: `context.WithChancel()` `context.WithTimeOut()` `context.WithDeadLine()` // 绝对时间

`context.WithValue()` // 多个goroutine传递值



## 日志收集项目

### 为什么要自己写不用ELK?

ELK: 部署的时候麻烦每一个filebeat都需要配置一个配置文件

使用etcd来管理被收集的日志项

### 项目框架

![image-20210201152231377](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201152231377.png)



### 上节课项目进度

1. kafka: 消息队列
2. tailf: 从文件里读日志
3. go-ini: 解析配置文件

## 今日内容

## etcd

**etcd介绍**

![image-20210201154956235](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201154956235.png)

![image-20210201155433337](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201155433337.png)

![image-20210201160016437](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201160016437.png)

![image-20210201160045433](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201160045433.png)

​	![image-20210201160826950](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201160826950.png)

![image-20210201161246969](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201161246969.png)



**安装etcd**

去GitHub上搜 etcd-io 找到发布版本最新的3.4.0+都可以,下载解压

点 etcd.exe打开即可

![image-20210201162435991](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201162435991.png)

![image-20210201162335232](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201162335232.png)

![image-20210201180503558](D:\Go\src\chentianxiang.vip\studygo\day12\README.assets\image-20210201180503558.png)

## 使用etcd优化日志收集项目

### 从etcd中拉 配置

### 如何开启多个tailObj   (创建个管理者 mgr)

### 如何实现watch配置变更实现热更新 

`详情见day12里面的logagent代码`





# 本周任务

1, 学习方法,查资料 2. 学习思路 3. 人脉 

1. Raft 协议  
   1.  选举
   2. 日志复制机制
   3. 异常处理(脑裂)
   4. zookeeper的zad协议的区别
2. etcd的watch (热更新)
   1. etcd底层如何实现watch给客户端发通知(websocket)
3. 课上的代码写一遍





