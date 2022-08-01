package main

import "fmt"

type person struct {
	name, city string
	age        int
}

// 构造函数
func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// 方法 值类型接收者
func (p person) showPerson() {
	fmt.Printf("my name is %s\n", p.name)
}

// 指针类型接收者
func (p *person) setName(name string) {
	p.name = name
}

// 指针类型接收者
func (p person) setNameValue(name string) {
	p.name = name
}

// 结构体的匿名字段
type calss struct {
	string
	int
}

func main() {
	p1 := newPerson("lili", "济南", 20)
	fmt.Println(p1)
	p2 := *p1
	p2.name = "kity"
	fmt.Println(p2)
	p2.setName("tom")
	p2.showPerson()
	p2.setName("jemmy")
	p2.showPerson()

	c1 := calss{
		"c1",
		10,
	}
	fmt.Println(c1)
	fmt.Println(c1.string, c1.int)

}
