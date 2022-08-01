package main

import (
	"fmt"
)

func main() {
	//if 语句
	if age := 10; age > 5 {
		fmt.Println("lower")
	} else {
		fmt.Println("higher")
	}
	//for循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	j := 2
	for j < 5 {
		fmt.Println(j)
		j++
	}
	//无线循环
	//for {
	//	print(2)
	//}
	//for ranger
	for index, value := range "hello go!" {
		fmt.Printf("%d:%d\n", index, value)
	}
	//switch case
	typeKey := 1
	switch typeKey {
	case 1:
		fmt.Println("age 1")
	case 2:
		fmt.Println("age 2")
	default:
		fmt.Println("none")
	}
	//一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switch mem := 2; mem {
	case 1, 3:
		fmt.Println("奇数")
	case 2, 4:
		fmt.Println("偶数")
	}
	//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
	peo := 30
	switch {
	case peo < 25:
		fmt.Println("low")
	case peo > 25 && peo < 35:
		fmt.Println("low")
	case peo > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
	//goto 跳转指定标签
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 4 {
			goto LABEL5
		}
	}
LABEL5:
	fmt.Println("=========")
}
