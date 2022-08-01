package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string `json:"gender"`
	Name   string
}
type Class struct {
	Title    string
	Students []*Student
}
type ApiResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// json 序列化
func main() {
	fmt.Println("#############")
	c := &Class{
		"一年级",
		make([]*Student, 0, 20),
	}
	api := &ApiResult{
		"0000",
		"success",
	}
	apiData, err := json.Marshal(api)
	fmt.Printf("json:%s\n", apiData)
	fmt.Println(c)
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	c1 := &Class{}
	json.Unmarshal([]byte(data), c1)
	fmt.Println(c1)
}
