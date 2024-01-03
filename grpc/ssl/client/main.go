package main

import (
	"context"
	"fmt"
	"go_learn/grpc/ssl/pb/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/3 11:56
 */

func main() {
	//使用单文件证书传输
	cds, err := credentials.NewClientTLSFromFile("../server.pem", "grpc.xxcheng.cn")
	if err != nil {
		fmt.Println("NewClientTLSFromFile err")
		panic(err)
	}
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(cds))
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
