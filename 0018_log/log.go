package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[...log...]")
}
func main() {

	//log.Println("log")
	//log.Printf("hello：%s", "momo")
	//log.Fatalln("fatal")
	////log.Panicln("panic")
	//log.Println("hi log")

	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}
