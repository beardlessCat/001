package main

import (
	"fmt"
	"strings"
)

/*
*
map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
*/
func main() {
	a := map[string]string{}
	fmt.Println(a)
	m := make(map[string]string, 7)
	m["1"] = "1"
	m["2"] = "2"
	fmt.Println(m)
	//判断是否存在某个值
	v, ok := m["3"]
	fmt.Println(v, ok)
	//遍历
	for k, v := range m {
		fmt.Println(k + "=" + v)
	}
	//元素为map的切片
	mapSlice := make([]map[string]string, 7)
	fmt.Println(mapSlice)
	mapSlice[0] = make(map[string]string)
	mapSlice[0]["name"] = "lyj"
	mapSlice[0]["age"] = "28"
	mapSlice[0]["sex"] = "男"
	fmt.Println(mapSlice)
	//值为切片类型的map
	m2 := make(map[string][]string, 3)
	fmt.Println(m2)
	key := "001"
	cv, cok := m2[key]
	fmt.Println(cv, cok)
	stringss := append(cv, "0001", "002")
	m2[key] = stringss
	fmt.Println(m2)
	//删除
	delete(m2, "001")
	fmt.Println(m2)

	str := "how do you do"
	split := strings.Split(str, " ")
	result := make(map[string]int)
	for _, v := range split {
		_, ok := result[v]
		if ok {
			result[v] += 1
		} else {
			result[v] = 0
		}
	}
	fmt.Println(result)
}
