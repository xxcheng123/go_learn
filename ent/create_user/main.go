package main

import (
	"context"
	"ent-demo/ent"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "ent_admin:12345678@tcp(192.168.21.100:3306)/ent?parseTime=True")
	if err != nil {
		fmt.Println("connect database failed.")
		panic(err)
	}
	defer client.Close()
	result, err := client.User.Create().
		SetID(time.Now().UnixMilli()).
		SetUsername(fmt.Sprintf("xxcheng_%d", time.Now().Unix())).
		SetAge(99).
		Save(context.Background())
	fmt.Println(result, err)
}
