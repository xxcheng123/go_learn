package main

import (
	"fmt"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/16 16:40
 */

func main() {
	defer func() {
		fmt.Println("Ok")
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	defer func() {
		panic("come from defer")
	}()
	defer func() {
		panic("come from defer222")
	}()

	panic("my panic")

	defer func() {
		fmt.Println("Hello")
	}()
}
