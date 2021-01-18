## 仓库用于存放学习go的代码



具体的关于Go的笔记可以参考我的博客.

[博客园博客](https://www.cnblogs.com/zhen1996)

[作者的博客](http://chentianxiang.vip/)

越努力,越幸运!

### vsCode 快捷键设置

`Ctrl/Command+Shift+P`,按下图输入`>snippets`，选择命令并执行

然后在弹出的窗口点击选择`go`选项：

```go
// 常用快捷提示设置

	"println":{
		"prefix": "pln",
		"body":"fmt.Println(\"$0\")",
		"description": "println"
	},
	"printf":{
		"prefix": "plf",
		"body": "fmt.Printf(\"$0\")",
		"description": "printf"
	},
	"func main()":{
		"prefix": "main",
		"body": "func main(){$0}",
		"description": "main函数"
	},
	"if err":{
		"prefix": "iferr",
		"body": "if err != nil {\n    fmt.Printf(\"$0,err:%v\\n\" ,err)\n    return\n}",
		"description": "if err "
	},
```

