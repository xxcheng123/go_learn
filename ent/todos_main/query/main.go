package main

import (
	"context"
	"ent-demo/db"
	"ent-demo/ent/todo"
	"fmt"
)

func main() {
	client := db.Get()
	items, err := client.Todo.Query().All(context.Background())
	if err != nil {
		panic(err)
	}
	for index, item := range items {
		fmt.Printf("index:%d,id:%d,title:%s\n", index, item.ID, item.Title)
	}
	items, err = client.Todo.Query().Where(todo.IDEQ(5)).All(context.Background())
	if err != nil {
		panic(err)
	}
	for index, item := range items {
		fmt.Printf("index:%d,id:%d,title:%s\n", index, item.ID, item.Title)
	}
}
