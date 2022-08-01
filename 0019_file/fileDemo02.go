package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./0019_file/fileDemo02.go")
	if err != nil {
		fmt.Println("file open error:", err)
		return
	}
	defer file.Close()
	var temp = make([]byte, 512)
	n, err := file.Read(temp)
	if err == io.EOF {
		fmt.Println("file read complete!")
	}
	if err != nil {
		fmt.Println("file read error:", err)
	}
	fmt.Printf("%d字节数据被读取\n", n)
	fmt.Println(string(temp[:n]))
}
