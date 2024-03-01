package main

import (
	"context"
	"ent-demo/db"
)

func main() {
	client := db.Get()
	us, _ := client.User.Query().All(context.Background())
	for _, u := range us {
		u.Update().AddGroupIDs(1).Exec(context.Background())
	}
}
