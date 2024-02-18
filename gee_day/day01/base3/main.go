package main

import (
	"fmt"
	"gee"
	"net/http"
	"strconv"
	"time"
)

func main() {
	engine := gee.New()
	engine.GET("/time", func(w http.ResponseWriter, req *http.Request) {
		t := time.Now().Unix()
		w.Write([]byte(strconv.FormatInt(t, 10)))
	})
	engine.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})
	fmt.Println("running...")
	engine.Run(":8080")
}
