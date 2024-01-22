package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"go_learn/asynq/task"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/17 16:34
 */

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	const redisAddr = "xx.xxcheng.cn:6379"
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()
	tasks := []struct {
		Type    string
		To      string
		Content string
	}{
		{Type: task.TaskSendEmail, To: "xxcheng", Content: "hello world"},
		{Type: task.TaskSendSms, To: "JPC", Content: "This is an apple."},
	}
	for _, item := range tasks {
		var willTask *asynq.Task
		switch item.Type {
		case task.TaskSendEmail:
			willTask = task.NewSendEmailTask(item.To, item.Content)
		case task.TaskSendSms:
			willTask = task.NewSendSmsTask(item.To, item.Content)
		}
		if info, err := client.Enqueue(willTask); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(info)
		}
	}
}
