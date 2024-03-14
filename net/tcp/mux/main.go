package main

import (
	"net"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-14 18:08:18
 * @LastEditTime: 2024-03-14 18:20:30
 */

func main() {
	dialer := net.Dialer{
		LocalAddr: &net.TCPAddr{
			Port: 9987,
		},
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
			})
		},
	}
	conn, err := dialer.Dial("tcp", "www.qq.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go func() {
		dialer := net.Dialer{
			LocalAddr: &net.TCPAddr{
				Port: 9987,
			},
			Control: func(network, address string, c syscall.RawConn) error {
				return c.Control(func(fd uintptr) {
					syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
					syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
				})
			},
		}
		conn, err := dialer.Dial("tcp", "www.baidu.com:80")
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		if err != nil {
			panic(err)
		}
		conn.Write([]byte("hello"))
		time.Sleep(time.Second * 5)
		conn.Write([]byte("hello"))
	}()
	conn.Write([]byte("hello"))
	time.Sleep(time.Second * 5)
	conn.Write([]byte("hello"))

}
