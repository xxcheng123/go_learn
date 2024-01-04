package main

import (
	"fmt"
	"go_learn/rpc/protoc/pb/simple"
	"google.golang.org/protobuf/proto"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/4 14:39
 */

func main() {
	req := &simple.HelloReq{
		Title: "A",
		Power: 999,
		Ok:    false,
	}
	fmt.Println("---req---")
	fmt.Println(req)
	fmt.Println("---req---")
	bs, err := proto.Marshal(req)
	if err != nil {
		return
	}
	fmt.Println("---marshal---")
	fmt.Println(len(bs), bs)
	fmt.Println(string(bs))
	fmt.Println("---marshal---")
	fmt.Println("---convert---")
	bs[2] = 66
	req2 := new(simple.HelloReq)
	err = proto.Unmarshal(bs, req2)
	if err != nil {
		return
	}
	fmt.Printf("%+v", req2)
}
