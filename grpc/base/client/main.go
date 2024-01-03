package main

import (
	"context"
	"fmt"
	"go_learn/grpc/base/pb/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/3 11:56
 */

func main() {
	//使用不安全的方式传输
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := simple.NewSimpleServiceClient(conn)
	resp, err := client.Hello(context.Background(), &simple.HelloReq{
		Title: "a b c",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp:%s\n", resp.Reply)
}
