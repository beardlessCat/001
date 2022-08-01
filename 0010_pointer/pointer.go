package main

import "fmt"

/*
*每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
取变量指针的语法如下：
ptr := &v    // v的类型为T
*/
func main() {
	a := 100
	//获取指针 ptr := &v
	aPtr := &a
	fmt.Printf("%d\n", aPtr)
	fmt.Printf("%T\n", aPtr)
	fmt.Println(&aPtr)

	//通过指针获取值 *ptr
	b := *aPtr
	fmt.Println(b)

}
