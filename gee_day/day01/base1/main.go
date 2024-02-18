package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Unix()
		w.Write([]byte(strconv.FormatInt(t, 10)))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	fmt.Println("start running.")
	http.ListenAndServe(":8080", nil)
}
