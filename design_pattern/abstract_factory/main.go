package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 16:17
 */

/**
抽象工厂方法
抽象工厂类：是一组创建产品的对象的抽象
实现工厂类
抽象类
实现类

抽象工厂有多个实现类，一个工厂可以创建多种实现类
*/

type Person interface {
	Say()
}
type Flower interface {
	Name()
}
type AbstractFactory interface {
	CreatePerson() Person
	CreateFlower() Flower
}
type ChinesePerson struct {
}

func (c *ChinesePerson) Say() {
	fmt.Println("中国人")
}

type ChineseFlower struct {
}

func (c *ChineseFlower) Name() {
	fmt.Println("中国花")
}

type ChineseFactory struct {
}

func (c *ChineseFactory) CreatePerson() Person {
	return &ChinesePerson{}
}
func (c *ChineseFactory) CreateFlower() Flower {
	return &ChineseFlower{}
}

type AmericanPerson struct {
}

func (c *AmericanPerson) Say() {
	fmt.Println("美国人")
}

type AmericanFlower struct {
}

func (c *AmericanFlower) Name() {
	fmt.Println("美国花")
}

type AmericanFactory struct {
}

func (c *AmericanFactory) CreatePerson() Person {
	return &AmericanPerson{}
}
func (c *AmericanFactory) CreateFlower() Flower {
	return &AmericanFlower{}
}

func main() {
	fc1 := &ChineseFactory{}
	fc2 := &AmericanFactory{}
	p1 := fc1.CreatePerson()
	f1 := fc1.CreateFlower()
	p1.Say()
	f1.Name()
	p2 := fc2.CreatePerson()
	f2 := fc2.CreateFlower()
	p2.Say()
	f2.Name()
}
