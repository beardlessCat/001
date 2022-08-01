package main

import "fmt"

/*
*
单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
虽然可以使用共享内存进行数据交换，但是共享内存在不同的 goroutine 中容易发生竞态问题。
为了保证数据交换的正确性，很多并发模型中必须使用互斥量对内存进行加锁，这种做法势必造成性能问题

如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。
channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，
总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，
也就是声明channel的时候需要为其指定元素类型。
*/
/**
var 变量名称 chan 元素类型

*/
type user struct {
	name string
	age  int
}

func cho(channel chan int) {
	value, ok := <-channel
	if !ok {
		fmt.Println("exception！")
	}
	fmt.Println(value)
}

func f3(channel chan int) {
	for v := range channel {
		fmt.Println(v)
	}
}
func main() {
	//有缓冲的通道
	//channel 定义
	var intCha chan int
	fmt.Println(intCha)
	//channel 初始化
	intCha = make(chan int, 2)
	fmt.Println(intCha)
	//发送 将一个值发送到通道中。
	intCha <- 10
	//接收 从一个通道中接收值。
	x := <-intCha
	fmt.Println(x)
	//关闭channel
	close(intCha)
	/* 无缓冲的通道 */

	//ch := make(chan int)
	//ch <- 10
	/**
	因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，
	否则会一直处于等待发送的阶段。同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。
	就像田径比赛中的4x100接力赛，想要完成交棒必须有一个能够接棒的运动员，否则只能等待。简单来说就是无缓冲的通道必须有至少一个接
	收方才能发送成功。
	上面的代码会阻塞在ch <- 10这一行代码形成死锁，
	其中一种可行的方法是创建一个 goroutine 去接收值，
	*/
	ch := make(chan int)
	go func(ch chan int) {
		ret := <-ch
		fmt.Println(ret)
	}(ch)
	ch <- 10
	fmt.Println("send success!")
	/**
	首先无缓冲通道ch上的发送操作会阻塞，直到另一个 goroutine 在该通道上执行接收操作，这时数字10才能发送成功，
	两个 goroutine 将继续执行。相反，如果接收操作先执行，接收方所在的 goroutine 将阻塞，直到 main goroutine 中向该通道发送数字10。
	使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为同步通道。
	*/
	//多返回值模式
	channel := make(chan int, 2)
	channel <- 10
	channel <- 20
	cho(channel)
	cho(channel)
	close(channel)
	cho(channel)
	fmt.Println("=============")
	//for range 接收值
	channel2 := make(chan int, 2)
	channel2 <- 10
	channel2 <- 20
	close(channel2)
	f3(channel2)

}
