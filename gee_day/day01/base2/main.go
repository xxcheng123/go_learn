package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Engine struct {
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/time":
		t := time.Now().Unix()
		w.Write([]byte(strconv.FormatInt(t, 10)))
	case "/hello":
		w.Write([]byte("Hello World"))
	default:
		http.NotFound(w, r)
	}
}

func main() {
	fmt.Println("start running.")
	http.ListenAndServe(":8080", new(Engine))
}
