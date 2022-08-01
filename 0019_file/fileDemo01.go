package main

import (
	"fmt"
	"os"
)

/*
*
os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件实例调用close()方法能够关闭文件。
*/
func main() {
	file, err := os.Open("./fileDemo01.go")
	if err != nil {
		fmt.Println("file read error:", err)
	}
	file.Close()
	fmt.Println("")
}
