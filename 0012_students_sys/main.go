package main

import (
	"fmt"
	"os"
)

var i = 0

// Student 学生
type Student struct {
	id, age int
	name    string
}

// System 管理系统
type System struct {
	students []*Student
}

// 学生构造函数
func newStudent(id, age int, name string) *Student {
	if id == 0 {
		i++
		id = i
	}
	return &Student{
		id,
		age,
		name,
	}
}

// 管理系统构造函数
func newSystem() *System {
	return &System{
		make([]*Student, 0, 100),
	}
}

// 添加学生
func (s *System) addStu(student *Student) {
	stus := append(s.students, student)
	s.students = stus
}

// 编辑学生
func (s *System) editStu(student *Student) {
	for index, stu := range s.students {
		if stu.id == student.id {
			s.students[index] = student
			return
		}
	}
}

// 删除学生
func (s *System) delStu(id int) {
	j := 0
	for index, stu := range s.students {
		if stu.id != id {
			s.students[j] = s.students[index]
			j++
		}
	}
	s.students = s.students[:j]
}

// 展示学生
func (s *System) showStus() {
	for _, stu := range s.students {
		fmt.Println(*stu)
	}
}

func main() {
	sm := newSystem()
	for {
		fmt.Printf("##############  学生管理系统  ##################\n")
		fmt.Printf("请输入您的操作类型：\n1：新增学生信息\n2：修改学生信息\n3：删除学生信息\n4：查询所有学生信息\n5：退出系统\n")
		var doType int
		_, err := fmt.Scan(&doType)
		if err != nil {
			fmt.Println("输入信息有误！")
		}
		switch doType {
		case 1:
			var (
				age  int
				name string
			)
			fmt.Println("请输入学生姓名：")
			fmt.Scan(&name)
			fmt.Println("请输入学生年龄：")
			fmt.Scan(&age)
			stu := newStudent(0, age, name)
			sm.addStu(stu)
			sm.showStus()
			break
		case 2:
			var (
				age, id int
				name    string
			)
			fmt.Println("请输入学生编号：")
			fmt.Scan(&id)
			fmt.Println("请输入学生姓名：")
			fmt.Scan(&name)
			fmt.Println("请输入学生年龄：")
			fmt.Scan(&age)
			stu := newStudent(id, age, name)
			sm.editStu(stu)
			break
		case 3:
			var id int
			fmt.Println("请输入学生编号：")
			fmt.Scan(&id)
			sm.delStu(id)
			break
		case 4:
			sm.showStus()
			break
		case 5:
			fmt.Println("即将退出系统......")
			os.Exit(0)
			break
		default:
			fmt.Println("输入的信息有误")
		}
	}
}
