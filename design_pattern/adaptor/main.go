package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 18:24
 */

type ChargeV220IF interface {
	ChargeV220()
}
type Phone struct {
	Board string
}

func (p *Phone) ChargeV220(s *SocketV220) {
	fmt.Printf("%s手机正在充电...\n", p.Board)
	s.Use()
}

type Socket interface {
	Use()
}
type SocketV220 struct {
}

func (s *SocketV220) Use() {
	fmt.Println("220V插座使用中...")
}

type SocketV110 struct {
}

func (s *SocketV110) Use() {
	fmt.Println("110V插座使用中...")
}
