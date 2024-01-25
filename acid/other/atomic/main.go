package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/24 14:06
 */

func main() {
	var x int64
	var y int64
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				x++
				atomic.AddInt64(&y, 1)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("x:[%d],y:[%d]\n", x, y)
}
