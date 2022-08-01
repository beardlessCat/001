package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		str := string(buf[:n])
		fmt.Println("收到client端发来的数据：", str)
		conn.Write([]byte(str)) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:18080")
	if err != nil {
		fmt.Println("listen failed,error:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,error:", err)
		}
		go process(conn)
	}
}
