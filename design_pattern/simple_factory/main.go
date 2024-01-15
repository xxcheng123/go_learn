package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 15:38
 */

/***
简单工厂模式
工厂类
抽象父类
实现子类

一个工厂创建多个实现类
*/

type SayI interface {
	Say()
}
type Person interface {
	SayI
}

type Man struct {
	name string
}

func (m *Man) Say() {
	fmt.Printf("男人[%s]在说话\n", m.name)
}

type Woman struct {
	name string
}

func (m *Woman) Say() {
	fmt.Printf("女人[%s]在说话\n", m.name)
}

type Factory struct {
}

func (f Factory) NewFactory(name string, val string) Person {
	if name == "man" {
		return &Man{val}
	} else if name == "woman" {
		return &Woman{val}
	}
	return nil
}
func main() {
	p1, p2 := Factory{}.NewFactory("man", "无核化"), Factory{}.NewFactory("woman", "车厘子")
	p1.Say()
	p2.Say()
}
