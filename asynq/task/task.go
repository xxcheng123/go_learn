package task

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/17 16:24
 */

const TaskSendEmail = "Task::SendEmail"
const TaskSendSms = "Task::SendSms"

type TaskSendEmailPayload struct {
	To      string
	Content string
}

type TaskSendSmsPayload struct {
	To      string
	Content string
}

func NewSendEmailTask(to, content string) *asynq.Task {
	payload, _ := json.Marshal(TaskSendEmailPayload{To: to, Content: content})
	return asynq.NewTask(TaskSendEmail, payload)
}

func NewSendSmsTask(to, content string) *asynq.Task {
	payload, _ := json.Marshal(TaskSendSmsPayload{To: to, Content: content})
	return asynq.NewTask(TaskSendSms, payload)
}
