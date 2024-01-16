package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 18:18
 */
/**
装饰模式
原有的功能，添加新的功能
*/

type UserIF interface {
	Walk()
}
type User struct{}

func (u *User) Walk() {
	fmt.Println("走路...")
}

type Caller struct {
	User UserIF
}

func (c *Caller) Walk() {
	fmt.Println("打电话")
	c.User.Walk()
}
func main() {
	var u UserIF = &User{}
	u = &Caller{User: u}
	u.Walk()
}
