package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go_learn/grpc/cassl/pb/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
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
	// 使用ca
	crt, err := tls.LoadX509KeyPair("../crts/server.pem", "../crts/server.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair err")
		panic(err)
	}
	crtPool := x509.NewCertPool()
	ca, err := os.ReadFile("../crts/ca.pem")
	if err != nil {
		fmt.Println("ReadFile err")
		panic(err)
	}
	if ok := crtPool.AppendCertsFromPEM(ca); !ok {
		panic("failed to append ca certs")
	}
	cds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{
			crt,
		},
		ServerName: "grpc.xxcheng.cn",
		RootCAs:    crtPool,
	})
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
