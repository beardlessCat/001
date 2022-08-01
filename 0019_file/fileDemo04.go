package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./0019_file/1.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("file open error:", err)
	}
	defer file.Close()
	str := "hello file\n"
	//file.Write([]byte(str)) //写入切片
	//file.WriteString(str)   //写入字符串
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str) //将数据先写入缓存
	}
	writer.Flush()
}
