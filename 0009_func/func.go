package main

import (
	"fmt"
)

/*
*
Go语言中定义函数使用func关键字，具体格式如下：

	func 函数名(参数)(返回值){
		函数体
	}
*/
//无参无返回值
func vo() {
	fmt.Println("hello")
}

// 有参无返回值
func voIn(x, y string) {
	fmt.Println(x + y)
}

// 有参有返回值
func cal(x, y int) (t, s int) {
	t = x + y
	s = x - y
	return
}

// 可变参数
func mui(x int, arr ...int) (sum int) {
	sum = x
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 变量作用域 分为局部变量及全局变量
var age = 20 //全局变量

func or() {
	x := "123" //局部变量
	age = 18   //如果局部变量和全局变量重名，优先访问局部变量。
	fmt.Printf("age：%d\n", age)
	fmt.Println(x)
}

// 函数类型与变量
// 定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数作为参数
func cales(x, y int, op func(x, y int) int) int {
	return op(x, y)
}

// 函数作为返回值
func funRtn(typeStr string) func(x, y int) int {
	switch typeStr {
	case "+":
		return add
	default:
		return sub
	}
}

// defer Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，
// 将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
func out() {
	fmt.Println("start......")
	defer fmt.Println("1......")
	defer fmt.Println("2......")
	defer fmt.Println("3......")
	fmt.Println("end......")

}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 闭包
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	vo()
	voIn("hi", " go")
	t, s := cal(100, 20)
	fmt.Println(t, s)
	fmt.Println(mui(10, 10, 10))
	fmt.Println(mui(10, 10, 10, 10))
	or()
	var c calculation = sub
	c = add
	fmt.Println(c(1, 2))

	//函数作为参数
	i := cales(1, 2, add)
	fmt.Println(i)

	fmt.Println(funRtn("+")(1, 2))
	fmt.Println(funRtn("-")(1, 2))

	out()
	fmt.Println("###########################")
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
	fmt.Println("###########################")
	//匿名函数
	/**
	函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。
	匿名函数就是没有函数名的函数，匿名函数的定义格式如下：
	func(参数)(返回值){
	    函数体
	}
	*/
	//// 将匿名函数保存到变量
	addNm := func(x, y int) (int, int) {
		z := x + y
		c := x - y
		return z, c
	}
	var z, cvq = addNm(1, 2)
	fmt.Println(z, cvq)
	//闭包：指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}
