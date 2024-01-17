package main

import (
	"fmt"
	"testing"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/16 16:46
 */

func add(a, b int) int {
	fmt.Printf("%d+%d=%d\n", a, b, a+b)
	return a + b
}
func TestDefer(t *testing.T) {
	defer fmt.Println("Over")
	defer add(add(1, 2), add(9, 7))
	fmt.Println("Hello")
}
