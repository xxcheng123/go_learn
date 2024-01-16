package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 18:24
 */
/*
适配器模式
原本要不了的，那就加一层
*/

type ChargeV5IF interface {
	UseV5()
}
type SocketV5 struct {
}

func (s *SocketV5) UseV5() {
	fmt.Println("5V插座被使用中...")
}

type ChargeV220 interface {
	UseV220()
}
type SocketV220 struct {
}

func (s *SocketV220) UseV220() {
	fmt.Println("220V插座被使用中...")
}

type Phone struct {
	Brade string
}

// ChargeV5 只有5V可以充电的接口
func (p *Phone) ChargeV5(s ChargeV5IF) {
	fmt.Println("开始充电了...")
	s.UseV5()
}

//	AdaptorV5 设计一个适配器，支持220V转到5V
//
// 那就要支持一个5V的接口
type AdaptorV5 struct {
	SocketV220 *SocketV220
}

func (a *AdaptorV5) UseV5() {
	fmt.Println("适配器转换5V...")
	a.SocketV220.UseV220()
}

func main() {
	phone := Phone{Brade: "小米"}
	adaptor := AdaptorV5{SocketV220: &SocketV220{}}
	phone.ChargeV5(&adaptor)
}
