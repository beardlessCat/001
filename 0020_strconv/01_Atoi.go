package main

import (
	"fmt"
	"strconv"
)

/*
*
Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下。
func Atoi(s string) (i int, err error)
*/
func main() {
	s1 := "12"
	atoi, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("convert err:", err)
	} else {
		fmt.Printf("type：%T,value：%d", atoi, atoi)
	}

}
