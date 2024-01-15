package main

import (
	"fmt"
	"sync"
	"time"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/15 17:40
 */

// 单例模式

var wg sync.WaitGroup

type Timer struct {
}

func (t *Timer) Print() {
	fmt.Printf("current time is %v\n", time.Now())
}

var globalTimer Timer

func init() {
	wg.Add(1)
	globalTimer = Timer{}
	wg.Done()
}

func main() {
	wg.Wait()
	globalTimer.Print()
	globalTimer.Print()
}
