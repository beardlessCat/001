package main

import (
	"fmt"
	"math"
	"strings"
)

/*
*
Go语言中有丰富的数据类型，除了基本的整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）等。
*/
func main() {
	/* 整型 */
	a := 10
	fmt.Printf("%d \n", a) //十进制
	fmt.Printf("%b \n", a) //二进制
	b := 077               // 八进制  以0开头
	fmt.Printf("%o\n", b)  //八进制
	c := 0xfff             // 十六进制  以0x开头
	fmt.Printf("%x\n", c)  //十六进制
	/*浮点数 Go语言支持两种浮点型数：float32和float64*/
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.10f\n", math.Pi)
	/*复数*/
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
	/*布尔值  Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。*/
	var b1 bool
	fmt.Println(b1)
	/*字符串*/
	/**
	Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
	Go 语言里的字符串的内部实现使用UTF-8编码。字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符
	*/
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	/*多行字符串*/
	line := `
	第一行
	第二行
	第三行
	`
	fmt.Println(line)
	/*字符串的常用操作*/
	//求长度
	length := len(line)
	fmt.Println(length)
	//字符串拼接
	fmt.Println(line + "第四行")
	//字符串分割
	str := "123@345@789"
	split := strings.Split(str, "@")
	fmt.Println(split)
	//是否包含
	contains := strings.Contains(str, "1")
	fmt.Println(contains)
	//前缀
	prefix := strings.HasPrefix(str, "123")
	suffix := strings.HasSuffix(str, "789")
	fmt.Println(prefix, suffix)
	//字符串出现位置
	index := strings.Index(str, "@")
	indexL := strings.LastIndex(str, "@")
	fmt.Println(index)
	fmt.Println(indexL)
	//join操作
	balance := []string{"1", "2", "3", "4", "5"}
	join := strings.Join(balance, ",")
	fmt.Println(join)
	/*字符：组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来：
	Go 语言的字符有以下两种:
	uint8类型，或者叫 byte 型，代表一个ASCII码字符。
	rune类型，代表一个 UTF-8字符。
	*/
	s := "hello沙河"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	for _, v := range s {
		fmt.Printf("%v(%c) ", v, v)
	}
	fmt.Println()

	//修改字符串
	/**
	要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	*/
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))
	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
