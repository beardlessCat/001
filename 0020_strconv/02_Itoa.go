package main

import (
	"fmt"
	"strconv"
)

/*
*Itoa()函数用于将int类型数据转换为对应的字符串表示
 */
func main() {
	num := 10
	itoa := strconv.Itoa(num)
	fmt.Printf("type:%T,value:%s", itoa, itoa)
}
