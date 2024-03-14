package main

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-14 11:41:26
 * @LastEditTime: 2024-03-14 15:22:16
 */

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	// 访问模型
	// 权限适配器
	e, err := casbin.NewEnforcer("../model.conf", "../policy.csv")
	if err != nil {
		panic(err)
	}
	{
		sub := "user::common"
		obj := "/user/info"
		act := "get"
		ok, err := e.Enforce(sub, obj, act)
		if err != nil {
			panic(err)
		}
		if ok {
			fmt.Println("pass")
		} else {
			fmt.Println("forbit")
		}
	}
	{
		sub := "user::common"
		obj := "/user/info"
		act := "post"
		ok, err := e.Enforce(sub, obj, act)
		if err != nil {
			panic(err)
		}
		if ok {
			fmt.Println("pass")
		} else {
			fmt.Println("forbit")
		}
	}
	{
		sub := "user::common"
		obj := "/user/update"
		act := "delete"
		ok, err := e.Enforce(sub, obj, act)
		if err != nil {
			panic(err)
		}
		if ok {
			fmt.Println("pass")
		} else {
			fmt.Println("forbit")
		}
	}
	{
		sub := "c"
		obj := "/a"
		act := "get"
		ok, err := e.Enforce(sub, obj, act)
		if err != nil {
			panic(err)
		}
		if ok {
			fmt.Println("pass")
		} else {
			fmt.Println("forbit")
		}
	}
}
