package main

import (
	"fmt"
	"os"
)

// 学生管理系统

// 初始化管理者
var smr studentMgr // 声明一个全局的变量学生管理smr

// 菜单函数
func showMenu() {
	fmt.Println("welcome sms!")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5. 退出
	`)
}

func main() {
	// 管理者初始化  修改的全局的那个变量
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		// 等待用户输入选项
		fmt.Print("请输入要执行的序号:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是:", choice)
		// ?
		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:
			smr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚")
		}

	}

}
