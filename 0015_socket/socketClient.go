package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:18080")
	if err != nil {
		fmt.Println("client connected error", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		if err != nil {
			fmt.Println("read error:", err)
		}
		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("write error:", err)
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("readError:", err)
		}
		fmt.Println("收到服务器发来的消息：", string(buf[:n]))
	}
}
