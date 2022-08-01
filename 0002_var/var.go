package main

import "fmt"

/*
Go语言的变量声明格式为：
var 变量名 变量类型，举个例子：
var name string
var age int
var isOk bool
*/
func main() {
	//变量声明
	////批量声明
	var (
		a string
		b int
		c bool
		d float64
	)
	////变量初始化 var 变量名 类型 = 表达式
	var sex string = "女"
	////单次初始化多个
	var hobby, score = "篮球", 10
	////类型推导
	var chatType = "微信"
	var days = 10
	//短变量声明 在函数内部，可以使用更简略的 := 方式声明并初始化变量。
	m := "day"
	n := 10
	//匿名变量

	//常量
	const (
		pi = 3.1415
		e
	)
	//iota 是go语言的常量计数器，只能在常量的表达式中使用。
	const (
		age1 = iota
		age2
		_
		age4
	)
	//多个iota定义在一行
	//const (
	//	a, b = iota + 1, iota + 2 //1,2
	//	c, d                      //2,3
	//	e, f                      //3,4
	//)
	fmt.Printf("a:%s,b:%d,c:%t,d:%g\n", a, b, c, d)
	fmt.Printf("sex:%s\n", sex)
	fmt.Printf("hobby:%s,score:%d\n", hobby, score)
	fmt.Printf("chatType:%s\n", chatType)
	fmt.Printf("chatType:%d\n", days)
	fmt.Printf("m:%s,n：%d\n", m, n)
	fmt.Printf("age1:%d\n,age2:%d\n,age4:%d\n", age1, age2, age4)

}
