package main

import "fmt"

// panic 和 recover

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		// recover 尽量少用!!
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B") // 程序崩了
	fmt.Println("v")    // 永远执行不到
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
