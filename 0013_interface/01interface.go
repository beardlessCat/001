package main

import "fmt"

/*
在Go语言中接口（interface）是一种类型，一种抽象的类型。相较于之前章节中讲到的那些具体类型（字符串、切片、结构体等）更注重“我是谁”，
接口类型更注重“我能做什么”的问题。接口类型就像是一种约定——概括了一种类型应该具备哪些方法，在Go语言中提倡使用面向接口的编程方式实现解耦。
*/
type Singer interface {
	Singer()
}

type Bird struct {
	name string
}

func (b Bird) Singer() {
	fmt.Printf("%s 嘤嘤嘤\n", b.name)
}

type Dog struct {
	string
}

func (d Dog) Singer() {
	fmt.Printf("%s 汪汪汪\n", d.string)
}

func letSing(singer Singer) {
	singer.Singer()
}

type Zfb struct {
}

type Pay interface {
	pay()
}

func (z *Zfb) pay() {
	fmt.Println("pay by ZFB")
}

type Wechat struct {
}

func (w Wechat) pay() {
	fmt.Println("pay by WeChat")
}
func checkOut(pay Pay) {
	pay.pay()
}
func main() {
	b := Bird{"小鸟"}
	letSing(b)
	d := Dog{"旺财"}
	letSing(d)

	checkOut(&Zfb{})
	checkOut(&Wechat{})

}
