package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上go关键字就可以创建一个 goroutine ，
// 从而让该函数或方法在新创建的 goroutine 中执行。
// go f()  // 创建一个新的 goroutine 运行函数
/**
匿名函数也支持使用go关键字创建 goroutine 去执行。

go func(){
  // ...
}()

*/
//类似于java的countDownLatch
var wg = sync.WaitGroup{}

func hello(obj interface{}) {
	fmt.Println("hello!", obj)
	wg.Done()
}
func main() {
	/**
	Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。默认值是机器上的 CPU 核心数。
	例如在一个 8 核心的机器上，GOMAXPROCS 默认为 8。Go语言中可以通过runtime.GOMAXPROCS函数设置当前程序并发时占用的 CPU逻辑核心数。
	（Go1.5版本之前，默认使用的是单核心执行。Go1.5 版本之后，默认使用全部的CPU 逻辑核心数。）
	*/
	runtime.GOMAXPROCS(2)
	wg.Add(1)
	go hello("小熊猫") // 启动另外一个goroutine去执行hello函数
	fmt.Println("你好")
	wg.Wait()

	//启动多个goroutine
	/**
	操作系统的线程一般都有固定的栈内存（通常为2MB）,而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB），
	所以在 Go 语言中一次创建数万个 goroutine 也是可能的。并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小，
	Go 的 runtime 会自动为 goroutine 分配合适的栈空间。
	*/
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	wg.Wait()
	//go 使用匿名函数
	wg.Add(100)
	for i := 0; i < 100; i++ {
		//闭包问题
		//go func() {
		//	fmt.Println("hello", i)
		//	wg.Done()
		//}()
		go func(i int) {
			fmt.Println("hello", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
