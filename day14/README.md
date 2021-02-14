



# day14 笔记

## 一. gin入门 (代码都截图了,,,自己敲一下,这个老师真懒)

### 1.安装

​	![image-20210212120426896](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212120426896.png)

```bash
go get -u -v github.com/gin-gonic/gin   // 创建个go mod  然后导入包下载
```



在github上搜gin即可

### 2. 使用

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin的helloWorld

func main() {

	//1. 创建路由
	r := gin.Default()
	// 2. 绑定路由规则,执行的函数
	// gin.Context, 封装饿了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World")
	})
	// 3. 监听端口,默认在8080
	r.Run(":8000")
}
```



## 二. gin路由

![image-20210212123959893](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212123959893.png)

![](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212142720757.png)

![image-20210212142739305](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212142739305.png)![image-20210212150155012](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212150155012.png)

**后台**

![image-20210212150219356](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212150219356.png)

**前端**

![image-20210212150513480](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212150513480.png)

**结果**:

![image-20210212150540695](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212150540695.png)

![image-20210212151433542](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212151433542.png)

**后台**

![image-20210212151423746](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212151423746.png)

**前端**

![image-20210212151458996](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212151458996.png)

![image-20210212151521676](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212151521676.png)

![image-20210212154451888](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212154451888.png)

![image-20210212154440030](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212154440030.png)

### 8. routes group 

```go
func main() {
	r := gin.Default()
	// 路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}

	v2 := r.Group("/v1")
	{
		v2.GET("/login", login)
		v2.GET("submit", submit)
	}
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
  }
}
```

![image-20210212155927390](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212155927390.png)

![image-20210212160458086](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212160458086.png)



![image-20210212190437124](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212190437124.png)





![image-20210212190224988](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212190224988.png)

![image-20210212190402620](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212190402620.png)

![image-20210212190317862](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212190317862.png)

![image-20210212190456349](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212190456349.png)



![image-20210212191154690](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212191154690.png)

![image-20210212191146653](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212191146653.png)

**html**

![image-20210212191219196](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212191219196.png)



![image-20210212191712667](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212191712667.png)

![image-20210212191707135](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212191707135.png)



**操作:**

![image-20210212191753640](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212191753640.png)





![image-20210212195154948](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212195154948.png)

![image-20210212195217332](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212195217332.png)



![image-20210212195245517](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212195245517.png)

![image-20210212195140756](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210212195140756.png)



![image-20210213150603650](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213150603650.png)

![image-20210213150548771](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213150548771.png)

![image-20210213150626558](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213150626558.png)

![image-20210213150943782](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213150943782.png)

![image-20210213151826955](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213151826955.png)

![image-20210213151815235](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213151815235.png)

![image-20210213153744192](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213153744192.png)

![image-20210213153729848](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213153729848.png)

![image-20210213153716882](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213153716882.png)

![image-20210213154127513](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213154127513.png)

![image-20210213154352196](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213154352196.png)

![image-20210213160358262](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213160358262.png)

![image-20210213160253311](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213160253311.png)

![image-20210213160245522](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213160245522.png)

![image-20210213160340669](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213160340669.png)



![image-20210213223006872](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213223006872.png)

![image-20210213224506864](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213224506864.png)

![image-20210213224653200](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213224653200.png)



![image-20210213224548088](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213224548088.png)

![image-20210213224607728](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213224607728.png)

![image-20210213224632464](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210213224632464.png)



 ![image-20210214180540495](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214180540495.png)



![image-20210214180501542](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214180501542.png)

![image-20210214180524986](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214180524986.png)

![image-20210214180811014](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214180811014.png)

![image-20210214182643352](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214182643352.png)



![image-20210214182637100](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214182637100.png)

![image-20210214185527010](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214185527010.png)

![image-20210214185735449](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214185735449.png)

![image-20210214190617860](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214190617860.png)

![image-20210214190713441](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214190713441.png)

![image-20210214190914221](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214190914221.png)

![image-20210214191405613](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214191405613.png)

![image-20210214191640963](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214191640963.png)





![image-20210214202413055](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214202413055.png)

![image-20210214202455124](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214202455124.png)![image-20210214202518106](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214202518106.png)

![image-20210214202543677](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214202543677.png)

![image-20210214202552402](D:\Go\src\chentianxiang.vip\studygo\day14\README.assets\image-20210214202552402.png)

```go
增删改查!!  图书管理系统
用 数据库 写,还有那个gin写!
页面给写好了
```

