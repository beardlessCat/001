package main

import "fmt"

/*
*
因为数组的长度是固定的并且数组长度属于类型的一部分，
*/
func main() {
	fmt.Println("+++++++++++++++++++++++++++")
	//切片的定义
	//var name []
	//name:表示变量名
	//T:表示切片中的元素类型
	var a []string
	var b []int
	var c = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//切片的长度和容量
	//长度是指切片的长度，容量是指切片所对应的数组的初始长度
	//切片表达式
	aa := [5]int{1, 2, 3, 4, 5}
	bb := aa[1:3]
	cc := aa[:]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", cc, len(cc), cap(cc))
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", bb, len(bb), cap(bb))
	//make()函数构造切片
	ints := make([]int, 8, 9)
	fmt.Println(ints)
	//切片的赋值拷贝
	ints[0] = 19
	fmt.Println(ints)
	//切片遍历  与数组一致
	//append 添加元素
	api := append(ints, 3, 5)
	fmt.Println(api)
	iss := append(api, bb...)
	fmt.Println(iss)

	//切片复制
	ce := make([]int, 8)
	ice := copy(ce, iss)
	fmt.Println(ice)
	fmt.Println(ce)

}
