package main

import (
	"context"
	"fmt"
	"go_learn/grpc/ssl/pb/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"time"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/3 11:46
 */

type SimpleServiceServer struct {
	simple.SimpleServiceServer
}

func (s *SimpleServiceServer) Hello(ctx context.Context, req *simple.HelloReq) (resp *simple.HelloResp, err error) {
	title := req.Title
	reply := fmt.Sprintf("receive msg:[%s],reply:%d", title, time.Now().Unix())
	fmt.Println(reply)
	return &simple.HelloResp{
		Reply: reply,
	}, nil
}

func main() {
	// 使用单文件证书
	cds, err := credentials.NewServerTLSFromFile("../server.pem", "../server.key")
	if err != nil {
		fmt.Println("NewServerTLSFromFile err")
		panic(err)
	}
	s := grpc.NewServer(grpc.Creds(cds))
	service := new(SimpleServiceServer)
	simple.RegisterSimpleServiceServer(s, service)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("server working...")
	err = s.Serve(listener)
	defer s.Stop()
	if err != nil {
		panic(err)
	}
}
