package main

import "fmt"

// 嵌套结构体
type Address struct {
	value string
	id    int
}

type School struct {
	name string
	age  int
	addr Address
}

// 嵌套结构体匿名
type Company struct {
	name string
	Address
}

//结构体的“继承”

type Animal struct {
	name string
}

type Dog struct {
	Feet int8
	*Animal
}

func (a *Animal) move() {
	fmt.Printf("%s is moving\n", a.name)
}
func (d *Dog) bark() {
	fmt.Printf("%s wang ! wang ! wang !\n", d.name)
}

func main() {
	fmt.Println("###################")
	s1 := School{
		"第一中学",
		20,
		Address{
			"济南市",
			2,
		},
	}
	fmt.Println(s1)
	var compamy Company
	compamy.name = "alibbaa"
	compamy.Address.id = 2 // 匿名字段默认使用类型名作为字段名
	compamy.value = "杭州"   // 匿名字段可以省略
	fmt.Println(compamy)

	dog := Dog{
		10,
		&Animal{
			"旺财",
		},
	}
	dog.bark()
	dog.move()
}
