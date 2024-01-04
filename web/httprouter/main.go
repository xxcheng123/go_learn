package main

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/4 9:17
 */

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"net/http"
)

type Resp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

var handler = func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println(params)
	resp := &Resp{
		Code:    200,
		Message: "success",
	}
	bs, err := json.Marshal(resp)
	if err != nil {
		err = errors.Wrap(err, "json marshal error")
		fmt.Printf("%+v\n", err)
		_, _ = writer.Write([]byte("error"))
	} else {
		_, err = writer.Write(bs)
	}
}

func main() {
	r := httprouter.New()
	r.GET("/hello", handler)
	r.GET("/user/:id/detail", handler)
	r.GET("/admin/*xx", handler)
	fmt.Println("running...")
	err := http.ListenAndServe(":1234", r)
	if err != nil {
		panic(err)
	}
}
