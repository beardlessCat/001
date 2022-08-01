package main

import (
	"fmt"
	"sort"
)

func main() {
	//数组定义
	var a [2]int
	fmt.Println(a)
	//数组出初始化
	//初始化数组时可以使用初始化列表来设置数组元素的值。
	a = [2]int{1, 2}
	fmt.Println(a)
	//按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度
	b := [...]string{"1", "2", "3"}
	fmt.Println(b)
	//我们还可以使用指定索引值的方式来初始化数组，例如:
	c := [...]string{1: "x", 3: "k"}
	fmt.Println(c)
	/*数组遍历*/
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
	for i, v := range c {
		fmt.Println(i, v)
	}

	var as = make([]string, 5, 10)
	fmt.Println(as)

	for i := 0; i < 10; i++ {
		as = append(as, fmt.Sprintf("%v", i))
	}
	fmt.Println(as)
	var asort = [...]int{3, 7, 8, 9, 1}
	fmt.Println(asort)

	sort.Ints(asort[:])
	fmt.Println(asort)
}
