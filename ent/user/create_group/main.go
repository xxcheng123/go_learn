package main

import (
	"context"
	"ent-demo/db"
)

func main() {
	client := db.Get()
	client.Group.Create().SetName("hello").Save(context.Background())
}
