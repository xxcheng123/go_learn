package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 16:07
 */

/**
工厂方法模式
工厂抽象类
工厂实现类
对象抽象类
对象实现类

多个工厂创建多个实现类
*/

type SayI interface {
	Say()
}
type Person interface {
	SayI
}
type AbstractFactory interface {
	Create()
}
type Man struct {
}

func (m *Man) Say() {
	fmt.Println("hello, I am a man")
}

type ManFactory struct {
}

func (m ManFactory) Create() Person {
	return &Man{}
}
func main() {
	manFactory := ManFactory{}
	m := manFactory.Create()
	m.Say()
}
