package main

import (
	"context"
	"fmt"
	"go_learn/grpc/interceptor/pb/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	return &simple.HelloResp{
		Reply: reply,
	}, nil
}

func main() {
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()),
		// 使用拦截器
		grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			fmt.Println("----------------")
			fmt.Println(ctx)
			fmt.Println(req)
			fmt.Println(info)
			fmt.Println(handler)
			fmt.Println("----------------")
			return handler(ctx, req)
		}),
		grpc.StreamInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return handler(srv, ss)
		}),
	)
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
