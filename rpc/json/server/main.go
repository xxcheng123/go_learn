package main

import (
	"fmt"
	"go_learn/rpc/base/common"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/3 15:44
 */

type HelloService struct {
	User string
}

func (h *HelloService) Hello(request string, reply *string) error {
	fmt.Printf("receive msg:%s\n", request)
	*reply = fmt.Sprintf("hello,%s", h.User)
	return nil
}
func (h *HelloService) Ping(request int, reply *int) error {
	fmt.Printf("receive msg:%d\n", request)
	*reply = 111
	return nil
}
func (h *HelloService) Info(request common.User, reply *common.User) error {
	fmt.Println(request)
	*reply = common.User{
		Username: "服务器：" + request.Username,
		Age:      request.Age + 1,
	}
	return nil
}

func main() {
	rpcServer := rpc.NewServer()
	err := rpcServer.Register(&HelloService{
		User: "xxcheng",
	})
	if err != nil {
		fmt.Println("rpc register error")
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("net listen error")
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener Accept error")
			panic(err)
		}
		go func() {
			//rpc.ServeConn(conn)
			rpcServer.ServeCodec(jsonrpc.NewServerCodec(conn))
		}()
	}
}
