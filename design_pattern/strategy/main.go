package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/16 10:45
 */

type Eater interface {
	Eat()
}

type Gentleman struct {
}

func (g *Gentleman) Eat() {
	fmt.Println("The gentleman takes out his fork and eats gracefully.")
}

type Barbarian struct {
}

func (b *Barbarian) Eat() {
	fmt.Println("The barbarian takes out his sword and slams it into his chest.")
}

type Context struct {
	Person Eater
}

func main() {
	g := &Gentleman{}
	b := &Barbarian{}
	c1 := &Context{g}
	c2 := &Context{b}
	c1.Person.Eat()
	c2.Person.Eat()
}
