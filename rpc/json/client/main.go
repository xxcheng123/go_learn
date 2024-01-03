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
* @Date: 2024/1/3 15:50
 */

func main() {
	var a int8
	fmt.Println(a & 0x7f)
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	//client := rpc.NewClient(conn)
	var reply string
	if err = client.Call("HelloService.Hello", "hello", &reply); err != nil {
		panic(err)
	}
	fmt.Printf("reply:%s\n", reply)
	var reply2 int
	if err = client.Call("HelloService.Ping", 888, &reply2); err != nil {
		panic(err)
	}
	fmt.Printf("reply:%d\n", reply2)

	var reply3 common.User
	if err = client.Call("HelloService.Info", common.User{
		Username: "xxcheng",
		Age:      90,
	}, &reply3); err != nil {
		panic(err)
	}
	fmt.Printf("reply:%v\n", reply3)
}
