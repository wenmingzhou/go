package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student //变量声明
)

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s \n", k, v.name)
	}
}

//newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func addStudent() {
	//向allStudent 中添加一个新的学生
	// 1 创建一个新学生
	// 2 获取用户的输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号:")
	fmt.Scanln(&id)

	fmt.Print("请输入学生的姓名:")
	fmt.Scanln(&name)

	// 3 造学生
	newStu := newStudent(id, name)
	fmt.Println(newStu)
	// 4 追加到allStudent这个map中
	allStudent[id] = newStu

}

func delStudent() {
	var (
		deleteId int64
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&deleteId)
	delete(allStudent, deleteId)
}
func main() {
	allStudent = make(map[int64]*student, 50) //初始化(开辟内存空间)
	for {
		fmt.Println("欢迎管理学士管理系统!")
		fmt.Println(`
				1.查看所有学生
				2.新增学生
                3.删除学生
                4.退出程序
				`)
		fmt.Print("请输入要你干啥!")

		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项!", choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入了无效字符")

		}
	}
}
