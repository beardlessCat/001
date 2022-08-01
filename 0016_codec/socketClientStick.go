package main

import (
	"fmt"
	"goSpace/0016_codec/codec"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:18080")
	if err != nil {
		fmt.Println("connectedError:", err)
	}
	defer conn.Close()
	for i := 0; i < 100; i++ {
		bytes, err := codec.Encode(`Hello, Hello. How are you?`)
		if err != nil {
			break
		}
		conn.Write(bytes)
	}
}
