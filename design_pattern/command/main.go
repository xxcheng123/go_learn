package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/16 10:25
 */

// Invoker Invoker Interface
type Invoker interface {
	Execute()
}

// User Receiver
type User struct {
	Name string
}

func (u *User) Walk() {
	fmt.Printf("%s is walking.\n", u.Name)
}
func (u *User) Run() {
	fmt.Printf("%s is running.\n", u.Name)
}

// WalkInvoker Concrete  Command
type WalkInvoker struct {
	User *User
}

func (w *WalkInvoker) Execute() {
	w.User.Walk()
}

type RunInvoker struct {
	User *User
}

func (w *RunInvoker) Execute() {
	w.User.Run()
}

func main() {
	u := &User{Name: "Tom"}
	r := &RunInvoker{User: u}
	w := &WalkInvoker{User: u}
	r.Execute()
	w.Execute()
}
