package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	////stus在range遍历过程中 stu会对stus对应索引进行值拷贝
	////而在这里是将stu的地址传递给 m ， 因此 m 的值 指向同一片地址空间
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	fmt.Println(stus)
	for _, stu := range stus {
		fmt.Println(stu.name, &stu)
		m[stu.name] = &stu
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	arr := []int{
		1, 2, 3, 4, 5,
	}
	for _, v := range arr {
		fmt.Println(&v)
	}
}
