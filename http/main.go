package main

import (
	"fmt"
	"go_learn/http/http"
	"io"
)

// "go_learn/http/http"

func main() {
	resp, _ := http.Get("http://www.qq.com:80/")
	fmt.Printf("%+v\n", resp)
	bs, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bs))
}
