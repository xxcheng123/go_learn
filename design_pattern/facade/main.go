package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/16 9:34
 */
/**
想要使用一个对象，但是这个对象可能包含多个对象，这些对象可能需要被外部访问，使用外观模式，将对象的内部对象隐藏起来，提供一个统一的接口，提供统一的访问方法，将对象的内部对象隐藏起来，提供一个统一的接口，提供统一的访问方法，将对象的内部对象隐藏起来，提供一个统一
*/
var hasPower = false

type TV struct {
}

func (t *TV) Open() {
	if !hasPower {
		fmt.Println("没有电")
		return
	}
	fmt.Println("打开电视")
}
func (t *TV) Close() {
	fmt.Println("关闭电视")
}

type Power struct {
}

func (p *Power) On() {
	hasPower = true
	fmt.Println("打开电源")
}
func (p *Power) Off() {
	fmt.Println("关闭电源")
}

type Facade struct {
	power *Power
	tv    *TV
}

func (f *Facade) Start() {
	f.power.On()
	f.tv.Open()
}
func (f *Facade) Over() {
	f.tv.Close()
	f.power.Off()
}
func main() {
	tv := &TV{}
	power := &Power{}
	power.On()
	tv.Open()

	facade := &Facade{tv: tv, power: power}
	facade.Over()
}
