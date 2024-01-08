package main

import (
	"runtime"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/8 9:41
 */

func main() {
	// 1、主动调度
	runtime.Gosched()
	// 2、加锁、互斥被动调度
	//gopark
	// 3、正常调度
	runtime.Goexit()
	// 4、抢占调度
	//使用 monitor g 进入内核态发起系统调用
	// 本地->全局->其他本地
}
