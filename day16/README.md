# day16笔记

## 面试题

![image-20210220170229166](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220170229166.png)

```go
package main

type S struct {
	m string
}

func f()*S{
	return &S{"foo"}	// A
}

func main() {
	p := f()	// B
	print(p.m)
}
```

![image-20210220170607602](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220170607602.png)

``` go
package main

const N = 3

func main() {
	m := make(map[int]*int)
	for i:= 0; i< N; i++{
		m[i] = &i
	}
	for _, v := range m {
		print(*v)
	}
}
```

修改后 

```go
package main

const N = 3

func main() {
	m := make(map[int]int)
	for i:= 0; i< N; i++{
		m[i] = i
	}
	for _, v := range m {
		print(v)
	}
}
```

![image-20210220170648761](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220170648761.png)

```go
package main

import "sync"

const n = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
}

```

修改后

```go
package main

import "sync"

const n = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(n)
    // 之前是for的问题,突然进去了,电脑好了就结束了
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
}

```



![image-20210220174325527](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220174325527.png)

![image-20210220174352130](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220174352130.png)

```go
package main

import (
	"math/rand"
	"sync"
)

const n = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			m[rand.Int()] = rand.Int()
		}()
	}
	wg.Wait()
	println(len(m))

}
```

解决方法

```go
package main

import (
	"math/rand"
	"sync"
)

const n = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	// 加上锁,解决并发
	flag := sync.Mutex{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
            flag.Lock()
			defer wg.Done()
			m[rand.Int()] = rand.Int()
			flag.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))

}

```

![image-20210220175653159](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220175653159.png)

![image-20210220175838952](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220175838952.png)

make  引用类型 map chan   

new 值类型

![image-20210220175941418](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220175941418.png)

修改

![image-20210220180125363](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220180125363.png)

![image-20210220180835990](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220180835990.png)

![image-20210220181011246](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220181011246.png)

![image-20210220181016236](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220181016236.png)

![image-20210220181638121](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220181638121.png)





![image-20210220190543680](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220190543680.png)

![image-20210220190859984](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220190859984.png)

![image-20210220191033518](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220191033518.png)

![image-20210220191543298](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220191543298.png)

![image-20210220192017191](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220192017191.png)

![image-20210220192110810](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220192110810.png)

![image-20210220192132603](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220192132603.png)



![image-20210220192946020](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220192946020.png)

![image-20210220193216937](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193216937.png)![image-20210220193336155](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193336155.png)

![image-20210220193417804](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193417804.png)

![image-20210220193459577](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193459577.png)

![image-20210220193521809](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193521809.png)

![image-20210220193622719](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193622719.png)

![image-20210220193657234](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193657234.png)

![image-20210220193720123](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193720123.png)

![image-20210220193736365](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193736365.png)

![image-20210220193748868](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193748868.png)

![image-20210220193848707](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193848707.png)

![image-20210220193901929](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193901929.png)

![image-20210220193931691](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220193931691.png)

![image-20210220195834229](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220195834229.png)

user.proto

```go
// 版本号

syntax = "proto3";

// 指定包
package proto;

// 定义结构体
message UserRequest{
    // 定义用户名
    string name =1;
}

// 响应结构体
message UserResponse{
    int32 id =1;
    string name =2;
    int32 age =3;
    // repeated 修饰符是可变数组,go转切片
    repeated string hobby = 4;
}

// service定义方法
service UserInfoService{
    rpc GetUserInfo(UserRequest) returns (UserResponse){
    }
}
```

![image-20210220201900551](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220201900551.png)





![image-20210220202623990](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220202623990.png)

![image-20210220202810437](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220202810437.png)

![image-20210220202929532](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220202929532.png)

![](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220203009532.png)



![image-20210220203606734](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220203606734.png)



![image-20210220203620811](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220203620811.png)

![image-20210220203543642](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220203543642.png)

![image-20210220203755186](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220203755186.png)

![image-20210220203844093](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220203844093.png)

![image-20210220203909949](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220203909949.png)

![image-20210220204326459](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220204326459.png)

![image-20210220204413431](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220204413431.png)

![image-20210220204442837](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220204442837.png)

其他的和自己前面写的类似!  

![image-20210220204520179](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220204520179.png)



## go-micro 示例(由于没有包,go微服务框架而已,届时需要再看)



![image-20210220204538670](D:\Go\src\chentianxiang.vip\studygo\day16\README.assets\image-20210220204538670.png)

### 2021-02-20 20:49:58

到此,go微服务已经完结,,go 的学习也接近了尾声!! 经历了一个多月陆续学习,结束了go !

不得不说go确实很强大!新生语言确实厉害! 还是回归java!!!

go 适合服务器开发,写中间件! 

下面还有面试题没看,抽空看看! 投了三个简历,都查看,都没回!.......

苦逼等待初始成绩,在家学go学的都忘了老子考过研了!2021-02-20 20:52:03

努力,越努力,越幸运!  加油!

明天开始回归java!!!!毕竟有python,go 等基础!so easy

以后少看视频,多看文档!文档学习还是很快的!!!!!!!!!!!!!!

整理go笔记,和自己的博客!!加油! 

永远不要高估自己!

