package main

import "fmt"

// 有一个物件
// 1. 它保存了一些数据----> 结构的字段
// 2. 它的4个功能	 ----> 结构体的方法

// 造一个学生类型
type student struct {
	id   int64
	name string
}

// 造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudents() {
	// 从s.allStudent这个map中把所有的学生逐个拿出来
	for _, stu := range s.allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudents() {
	// 1.根据用户输入的内容创建一个新的学生
	var (
		stuID   int64
		stuName string
	)
	// 获取用户输入
	fmt.Print("请输入学号:")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名:")
	fmt.Scanln(&stuName)
	// 根据用户输入创建结构体对象
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功")
}

// 修改学生
func (s studentMgr) editStudents() {
	//1. 获取用户输入
	var stuID int64
	fmt.Print("亲输入学号")
	fmt.Scanln(&stuID)
	// 2. 展示该学号对应的学生信息,如果没有提示查无此人
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下:学号:%d, 姓名:%s\n", stuObj.id, stuObj.name)
	//3. 请输入修改后的学生名
	fmt.Print("亲输入学生的新名字:")
	var newName string
	fmt.Scanln(&newName)
	//4. 更新学生的姓名
	stuObj.name = newName
	s.allStudent[stuID] = stuObj // 更新map中的学生
}

// 删除学生
func (s studentMgr) deleteStudents() {
	// 1. 请用户输入要删除的学生id
	var stuID int64
	fmt.Print("请用户输入要删除的学生id")
	fmt.Scanln(&stuID)
	// 2. 去map去找有没有这个学生,如果没有提示查无此人
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	// 3. 有的话就删除, 如何从map中删除键值对
	delete(s.allStudent, stuID)
	fmt.Println("删除成功")
}
