package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9977")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		fmt.Println("accept...")
		conn, err := l.Accept()
		fmt.Println("receive...")
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func() {
			fmt.Println(conn.RemoteAddr().String())
		}()
		fmt.Println("next")
	}
}
