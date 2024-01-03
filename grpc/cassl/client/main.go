package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go_learn/grpc/cassl/pb/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/3 11:56
 */

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
