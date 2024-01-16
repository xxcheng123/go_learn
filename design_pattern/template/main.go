package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/16 9:46
 */
/**
模板方法
抽象类定义一组方法
实现类去实现
*/

type EatIF interface {
	Prepare()
	Eating()
	After()
}
type TaskOut struct {
}

func (t *TaskOut) Prepare() {
	fmt.Println("点外卖，到了下楼去拿")
}
func (t *TaskOut) Eating() {
	fmt.Println("吃外卖")
}
func (t *TaskOut) After() {
	fmt.Println("吃完了，把外卖丢了")
}

type Restaurant struct {
}

func (r *Restaurant) Prepare() {
	fmt.Println("去饭店下单")
}
func (r *Restaurant) Eating() {
	fmt.Println("上菜开吃")
}
func (r *Restaurant) After() {
	fmt.Println("吃完了，散散步")
}

func main() {
	var t EatIF = &TaskOut{}
	var t2 EatIF = &Restaurant{}
	t.Prepare()
	t.Eating()
	t.After()
	t2.Prepare()
	t2.Eating()
	t2.After()
}
