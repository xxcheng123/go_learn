package main

import (
	"context"
	"ent-demo/ent"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "ent_admin:12345678@tcp(192.168.21.100:3306)/ent?parseTime=True")
	if err != nil {
		fmt.Println("connect database failed.")
		panic(err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Println("create schema failed.")
		panic(err)
	}
}
