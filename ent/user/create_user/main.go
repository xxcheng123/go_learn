package main

import (
	"context"
	"ent-demo/db"
	"fmt"
)

func main() {
	client := db.Get()
	u, err := client.User.Create().SetUsername("xxcheng12345").AddCarIDs(1709273315573885, 1709273961594079).Save(context.Background())
	fmt.Println(u, err)
}
