package db

import (
	"ent-demo/ent"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var client *ent.Client

func init() {
	var err error
	client, err = ent.Open("mysql", "ent_admin:12345678@tcp(192.168.21.100:3306)/ent?parseTime=True")
	if err != nil {
		fmt.Println("connect database failed.")
		panic(err)
	}
}

func Close() {
	client.Close()
}
func Get() *ent.Client {
	return client
}
