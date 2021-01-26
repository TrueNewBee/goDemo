package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // 导入的是新版本
)

// go get github.com/go-redis/redis

// 最新版本的go-redis库的相关命令都需要传递context.Context参数，例如：
// 具体可以看博客 有点出入!2021-01-26 21:01:06

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func main() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}
	fmt.Println("连接成功")
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}
