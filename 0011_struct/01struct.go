package main

import (
	"fmt"
)

/*
*
Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
*/
type Person struct {
	name, city string
	age        int
}

type class struct {
	id, name string
	score    int
}

func main() {
	//自定义类型
	type MyInt int
	var x MyInt
	fmt.Printf("%T\n", x)
	//类型别名
	type mInt = int
	var y mInt
	fmt.Printf("%T\n", y)
	//结构体
	/**
	Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，
	这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，
	这种数据类型叫结构体，英文名称struct。 也就是我们可以通过struct来定义自己的类型了。
	Go语言中通过struct来实现面向对象。
	type 类型名 struct {
	    字段名 字段类型
	    字段名 字段类型
	    …
	}
	*/
	//结构体定义及实例化
	var p1 Person
	p1.name = "张三"
	p1.city = "李四"
	p1.age = 18
	fmt.Printf("%T:%v\n", p1, p1)
	p2 := Person{
		"小米",
		"北京",
		20,
	}
	fmt.Printf("%T:%v\n", p2, p2)

	//匿名结构体
	var user struct {
		id    string
		score int
	}
	user.id = "11"
	user.score = 20
	fmt.Printf("%T:%v\n", user, user)
	//创建指针类型结构体,我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
	p3 := new(Person)
	fmt.Printf("%T\n", p3)
	p3.name = "xiaohu" //p3.name = "xiaohu"其实在底层是(*p3).name = "xiaohu"，这是Go语言帮我们实现的语法糖。
	p3.city = "北京"
	p3.age = 30
	fmt.Printf("p3=%#v\n", p3) //p2=&main.person{name:"小王子", city:"上海", age:28}

	fmt.Println(*p3)
	//结构体初始化
	//没有初始化的结构体，其成员变量都是对应其类型的零值。
	var c1 class
	fmt.Println(c1)
	//使用键值对进行初始化
	c2 := class{
		id:    "1",
		name:  "c1",
		score: 10,
	}
	fmt.Println(c2)

	//也可以对结构体指针进行键值对初始化，例如：
	c3 := &class{
		id:    "2",
		name:  "c3",
		score: 10,
	}
	fmt.Println(c3)
	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	c4 := &class{
		name: "c3",
	}
	fmt.Println(c4)
	//使用值的列表初始化
	//1.必须初始化结构体的所有字段。
	//2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	//3.该方式不能和键值初始化方式混用。
	c5 := &class{
		"5",
		"c5",
		1,
	}
	fmt.Println(c5)

}
