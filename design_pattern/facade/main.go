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

type Walker struct {
}

func (w *Walker) Walk() {
	fmt.Println("走路...")
}

type Runner struct {
}

func (r *Runner) Run() {
	fmt.Println("跑步...")
}

type Swimmer struct {
}

func (s *Swimmer) Swim() {
	fmt.Println("游泳...")
}

type Facade struct {
	walker  *Walker
	runner  *Runner
	swimmer *Swimmer
}

func (f *Facade) All() {
	f.swimmer.Swim()
	f.walker.Walk()
	f.runner.Run()
}
func main() {
	swimmer := &Swimmer{}
	runner := &Runner{}
	walker := &Walker{}
	swimmer.Swim()
	runner.Run()
	walker.Walk()
	fmt.Println("---使用外观模式---")
	facade := &Facade{walker: walker, runner: runner, swimmer: swimmer}
	facade.All()
}
