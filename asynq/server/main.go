package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go_learn/asynq/task"
	"log"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/17 16:24
 */
type HandleSendEmail struct {
}

func (h *HandleSendEmail) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p *task.TaskSendEmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	fmt.Printf("Sending Email to User: %s, Content: %s\n", p.To, p.Content)
	// Email delivery code ...
	return nil
}
func HandleSendSms(ctx context.Context, t *asynq.Task) error {
	var p *task.TaskSendSmsPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	fmt.Printf("Sending SMS to User: %s, Content: %s\n", p.To, p.Content)
	return nil
}

func main() {
	const redisAddr = "xx.xxcheng.cn:6379"
	server := asynq.NewServer(asynq.RedisClientOpt{Addr: redisAddr}, asynq.Config{
		Concurrency: 5,
	})
	fmt.Println("start running.")
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TaskSendSms, HandleSendSms)
	mux.Handle(task.TaskSendEmail, &HandleSendEmail{})
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
