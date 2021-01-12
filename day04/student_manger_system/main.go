package main

import (
	"fmt"
	"os"
)

/*
	其实就是map的CURD  类似于Python版本,,,可以做成方法版本
	函数版本学生管理系统
	写一个系统能够查看\新增\删除学生
*/
// 创建一个map
var (
	// velue是int key是 student指针类型的map
	allStudent map[int64]*student // 变量声明
)

// student类型
type student struct {
	id   int64
	name string
}

// 构造函数 创建学生
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

// 展示所有学生
func showAllStudent() {
	// v是指针传入的指针
	for k, v := range allStudent {
		fmt.Printf("学号: %d 姓名:%s\n", k, v.name)
	}
}

// 添加学生
func addStudent() {
	// 向allStudent中添加一个新的学生
	// 1.创建一个学生
	// 1.1 获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名")
	fmt.Scanln(&name)
	// 1.2 造学生(调用student的构造函数)
	newStu := newStudent(id, name)
	// 2. 追加到alllStudent这个map中
	allStudent[id] = newStu
}

// 删除uxuesheng
func deleteStudent() {
	// 1. 请用户输入要删除的学生的序号
	var (
		deleteID int64
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&deleteID)
	// 2. 去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteID)
}

func main() {
	allStudent = make(map[int64]*student, 50) // 初始化开辟内存空间
	for {                                     // 增加一个for 相当于 while
		// 1.打印菜单
		fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出系统
	`)
		fmt.Print("请输入你要干啥")
		// 2.等到用户选择要什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		// 3. 执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出
		default:
			fmt.Println("滚~~")
		}
	}
}
