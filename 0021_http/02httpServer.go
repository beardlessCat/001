package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

type MyHandler struct {
}

func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.HandleFunc("/helli", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server start failed!")
		return
	} else {
		fmt.Println("server has started success!")
	}
	var handler MyHandler
	s := http.Server{
		Addr:           ":8888",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
