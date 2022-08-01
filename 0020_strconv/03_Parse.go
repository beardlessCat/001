package main

import (
	"fmt"
	"strconv"
)

/*
*
Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
*/
func main() {
	floatStr := "1"
	intStr := "10"
	boolStr := "false"
	boo, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("ParseBool faild")
	} else {
		fmt.Printf("type:%T", boo)
	}
	i, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		fmt.Println("ParseBool faild")
	} else {
		fmt.Printf("type:%T", i)
	}
	float, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("ParseBool faild")
	} else {
		fmt.Printf("type:%T", float)
	}
}
