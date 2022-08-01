package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com/")
	data := url.Values{}
	data.Set("name", "tom")
	if err != nil {
		fmt.Println("http get error:", err)
		return
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Println(string(all))
}
