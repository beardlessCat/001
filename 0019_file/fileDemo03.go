package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./0019_file/fileDemo03.go")
	if err != nil {
		fmt.Println("file open error:", err)
		return
	}
	var temp = make([]byte, 128)
	var content []byte
	for {
		n, readErr := file.Read(temp)
		if readErr == io.EOF {
			fmt.Println("file read complete!")
			break
		}
		if readErr != nil {
			fmt.Println("file read error!")
		}
		content = append(content, temp[:n]...)
	}
	fmt.Printf("文件读取长度：%d\n", len(content))
	fmt.Println(string(content))
}
