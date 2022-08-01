package main

import "fmt"

type User struct {
	no   string
	name string
	age  int
}

type SmsSystem struct {
	students map[string]*User
}

func NewSmsSystem() *SmsSystem {
	return &SmsSystem{
		make(map[string]*User, 10),
	}
}
func (s SmsSystem) initSys() {
	user1 := NewUser("1", "bob", 12)
	user2 := NewUser("2", "tom", 22)
	user3 := NewUser("3", "jim", 32)
	s.students[user1.name] = user1
	s.students[user2.name] = user2
	s.students[user3.name] = user3
}

func NewUser(no string, name string, age int) *User {
	return &User{no: no, name: name, age: age}
}

func (u User) showUser() {
	fmt.Printf("|%s|%s|%d|\n", u.no, u.name, u.age)
}

func main() {
	fmt.Println("================欢迎使用学生管理系统==============")
	smsSystem := NewSmsSystem()
	smsSystem.initSys()
	for {
		fmt.Println("1:增加学生信息")
		fmt.Println("2:修改学生信息")
		fmt.Println("3:查询学生信息")
		fmt.Println("4:删除学生信息")
		fmt.Println("5:展示学生信息")
		var delType int
		_, err := fmt.Scan(&delType)
		if err != nil {
			fmt.Println("发生错误！")
		}
		switch delType {
		case 1:
			var no, name string
			var age int
			fmt.Println("请输入学生编号：")
			fmt.Scan(&no)
			fmt.Println("请输入学生姓名：")
			fmt.Scan(&name)
			fmt.Println("请输入学生年龄：")
			fmt.Scan(&age)
			user := NewUser(no, name, age)
			smsSystem.students[no] = user
			break
		case 2:
			fmt.Println("请输入要修改的学生编号：")
			var no, name string
			var age int
			fmt.Scan(&no)
			user := smsSystem.students[no]
			fmt.Println("请输入学生年龄：")
			fmt.Scan(&name)
			fmt.Println("请输入学生年龄：")
			fmt.Scan(&age)
			user.no = no
			user.name = name
			break
		case 3:
			fmt.Println("请输入查询的学生编号：")
			var no string
			fmt.Scan(&no)
			user := smsSystem.students[no]
			user.showUser()
			break
		case 4:
			fmt.Println("请输入要删除的学生编号：")
			var no string
			fmt.Scan(&no)
			delete(smsSystem.students, no)
			break
		case 5:
			fmt.Println("学生信息如下:\n|学号|姓名|年龄|")
			for _, user := range smsSystem.students {
				user.showUser()
			}
			break
		default:
			fmt.Println("输入信息有误，请重新输入！")

		}

	}
}
