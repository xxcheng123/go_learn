package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 17:57
 */
/***
原本不完善的功能，加一个代理模式，帮你做其他事情
*/

type SwimmingIF interface {
	Swimming()
}
type Person struct {
	Name string
}

const Rain = 1
const Sunny = 2

type Environment struct {
	Department int
	Weather    int
}

func (p *Person) Swimming() {
	fmt.Printf("%s去游泳了\n", p.Name)
}

// 代理模式
// 满足开闭原则
// 额外的添加一些功能：检查什么什么的
func main() {
	p := Person{Name: "小米"}
	e := Environment{Department: 38, Weather: Sunny}
	if e.Department > 30 {
		fmt.Println("气温检查：天气允许")
		if e.Weather == Sunny {
			fmt.Println("天气检查：天气允许")
			p.Swimming()
		} else {
			fmt.Println("天气检查：天气不允许")
		}
	} else {
		fmt.Println("气温检查：不通过")
	}
	fmt.Println("-------下面是代理模式-------")
	(&SwimmingProxy{
		swimming:    &p,
		environment: e,
	}).Swimming()
}

type SwimmingProxy struct {
	swimming    SwimmingIF
	environment Environment
}

func (p *SwimmingProxy) Swimming() {
	fmt.Println("开始执行代理")
	if p.environment.Department > 30 {
		fmt.Println("气温检查：天气允许")
		if p.environment.Weather == Sunny {
			fmt.Println("天气检查：天气允许")
			p.swimming.Swimming()
		} else {
			fmt.Println("天气检查：天气不允许")
		}
	} else {
		fmt.Println("气温检查：不通过")
	}
	fmt.Println("结束执行代理")
}
