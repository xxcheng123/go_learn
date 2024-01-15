package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 18:18
 */

type User interface {
	GetName() string
	SetName(string)
}
type UserImpl struct {
	name string
}

func (u *UserImpl) GetName() string {
	return u.name
}
func (u *UserImpl) SetName(name string) {
	u.name = name
}

type UserDecorator struct {
	User User
}

func (u *UserDecorator) Walk() {
	fmt.Printf("%s出去走走\n", u.User.GetName())
}
func main() {
	var u User = &UserImpl{}
	u.SetName("大米")
	ud := UserDecorator{User: u}
	ud.Walk()
}
