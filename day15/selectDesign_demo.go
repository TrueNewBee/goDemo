package main

import "fmt"

// 通过 选项设计模式,实现结构体字段增加,方法不变

//结构体
type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	strOption4 string
	intOption1 int
	intOption2 int
	intOption3 int
}

// 定义一个函数类型变量
type Option func(opts *Options)

// 初始化结构体
func InitOptions1(opts ...Option) {
	options := &Options{}
	// 遍历 opts,得到每一个函数
	for _, opt := range opts {
		// 调用函数,在函数里,给传进去的对象赋值
		opt(options)
	}
	fmt.Printf("init options %#v\n", options)
}

// 定义具体给某个字段赋值的方法
func WithStrOption1(str string) Option {
	return func(opts *Options) {
		opts.strOption1 = str
	}
}
func WithStrOption2(str string) Option {
	return func(opts *Options) {
		opts.strOption2 = str
	}
}
func WithIntOption1(i int) Option {
	return func(opts *Options) {
		opts.intOption1 = i
	}
}

func main() {
	InitOptions1(WithIntOption1(1), WithStrOption1("hello"), WithStrOption2("go"))
}
