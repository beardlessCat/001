package main

import (
	"bufio"
	"fmt"
	"goSpace/0016_codec/codec"
	"net"
)

func processAccept(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := codec.Decode(reader)
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		if message == "" {
			fmt.Println("message is not full:")
			return
		}
		fmt.Println("收到client发来的数据：", message)
	}
}
func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:18080")
	if err != nil {
		fmt.Println("listen error:", err)
	}
	defer conn.Close()
	for {
		accept, err := conn.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
		}
		go processAccept(accept)
	}
}
