package main

import (
	"github.com/gin-gonic/gin"
)

// gin的helloWorld

func main() {
	//1. 创建路由
	// 默认使用了2个中间件 Logger(), Recover()
	r := gin.Default()
	//// 2. 绑定路由规则,执行的函数
	//// gin.Context, 封装饿了request和response
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World")
	//})
	//r.POST("/xxxPost", getting)
	//r.PUT("/xxxPut")
	//// 3. 监听端口,默认在8080
	//r.Run(":8000")

	// api 参数
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	c.String(http.StatusOK, name+ " is "+ action)
	//})
	//r.Run(":8000")

//	// 路由组
//	v1 := r.Group("/v1")
//	{
//		v1.GET("/login", login)
//		v1.GET("submit", submit)
//	}
//
//	v2 := r.Group("/v1")
//	{
//		v2.GET("/login", login)
//		v2.GET("submit", submit)
//	}
//}
//
//func login(c *gin.Context) {
//	name := c.DefaultQuery("name", "jack")
//	c.String(200, fmt.Sprintf("hello %s\n", name))
//}
//func submit(c *gin.Context) {
//	name := c.DefaultQuery("name", "jack")
//	c.String(200, fmt.Sprintf("hello %s\n", name))

	// 多相应方式


	r.Run(":8000")

}

