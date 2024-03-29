package main

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-14 18:00:55
 * @LastEditTime: 2024-03-14 18:07:39
 */

import (
	"fmt"
	"net"
	"time"
)

func main() {
	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			Port: 9816,
		},
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", "60.185.105.220:9977")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		fmt.Println("sending...")
		conn.Write([]byte("GET / HTTP/1.1\r\nHost: www.qq.com\r\n\r\n"))
		time.Sleep(1 * time.Second)
	}
}
