# day11课上笔记

# 今日内容

## 依赖管理go module

1.15默认开启GO111MODULE的

直接无脑安装,不需要配置工作区的环境变量

可以手动开启 `set GO111MODULE=on ` 

可以设置环境变量 `go env -w GO111MODULE=on`

更改 go get 的源 `go env -w GOPROXY=https://goproxy.cn,direct`

初始化init 需要 在项目下  `go mod init 该包的名字`  如`go mod init xxx`  这样go build后的 exe文件就是该名字
ps: **go buidl 之前,必须有mod文件, 先 执行上面步骤**

![image-20210127183720501](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210127183720501.png)

![image-20210127183713603](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210127183713603.png)

然后 `go get ` 下载文件就好了

![image-20210128221342596](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210128221342596.png)

## Context

## 服务端Agent开发

### 项目架构设计

### kafka和zookeeper

### tailf介绍