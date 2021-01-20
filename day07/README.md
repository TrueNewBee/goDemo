# day07课上笔记

# 内容回顾

## time

`2006-01-02 15:04:05:000`

### 时间类型

- `time.Time`:`time.now()`
- 时间戳:
  - `time.Now().Unix()`: 1970.1.1到现在的秒数
  - `time.Now().UnixNano()`:1970.1.1 到现在的纳米秒

### 时间间隔类型

- `time.Duration`: 时间间隔类型
  - `time.Second`

```go
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
```

## 时间操作

时间对象+/- 一个时间间隔对象

```go
// 时间间隔
fmt.Println(time.Second)
// now + 24小时
fmt.Println(now.Add(24 * time.Hour))	
// Sub 两个时间相减
nextYear, err := time.Parse("2006-01-02", "2021-08-04")
if err != nil {
    fmt.Printf("parse time failed, err:%v\n", err)
    return
}
d := nextYear.Sub(now)
```

after/befor

### 时间格式化

2006-01-02  15:04:05.000

### 定时器

```go
// 定时器	
timer := time.Tick(time.Second)
for t := range timer {
    fmt.Println(t) // 1秒钟执行一次
}
```

### 解析字符串格式的时间(时区)

```go
func f2() {
	now := time.Now() // 本地时间
	fmt.Println(now)
	// 明天这个时间
	// 按照指定格式取解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2021-01-17 16:31:29")
	// 按照东八区的时区和格式取解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed ,err: %v\n", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-17 16:31:29", loc)
	if err != nil {
		fmt.Printf("load loc failed ,err: %v\n", err)
		return
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}
```

## 日志库

time

文件操作

`runttime.Caller()`

之前学的内容整合

## 反射

 接口类型的变量底层是分为两个部分:动态类型和动态值

反射的应用:`json`等数据解析\ORM等工具

### 反射的两个方法

- `reflect.TypeOf()`
- `reflect.ValueOf()`



## ini解析作业

自己可以封装成 go ini   发到GitHub 开源项目 

### 面向GitHub 编程  233333

![image-20210119204310451](D:\Go\src\chentianxiang.vip\studygo\day07\README.assets\image-20210119204310451.png)



# 今日内容

并发

goutine

channel

sync

