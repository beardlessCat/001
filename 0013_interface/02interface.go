package main

import "fmt"

// Mover 定义一个接口类型
type Mover interface {
	Move()
}

// Dog 狗结构体类型
type Dogger struct{}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dogger) Move() {
	fmt.Println("狗会动")
}

// Cat 猫结构体类型
type Cat struct{}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}
func clear(washingMachine WashingMachine) {

}

// 接口组合

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
type WriteCloser interface {
	Writer
	Closer
}

// 空接口
// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 空接口作为map值
var ma = make(map[string]interface{})

func main() {
	var x Mover
	d1 := Dogger{}
	x = d1
	x.Move()
	x = &Dogger{}
	x.Move()

	var c1 = &Cat{} // c1是*Cat类型
	x = c1          // 可以将c1当成Mover类型
	x.Move()

	// 下面的代码无法通过编译
	//var c2 = Cat{} // c2是Cat类型
	//x = c2         // 不能将c2当成Mover类型

	clear(haier{})
}
