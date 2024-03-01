package main

import (
	"context"
	"ent-demo/db"
	"ent-demo/ent/user"
	"fmt"
)

func main() {
	client := db.Get()
	u, err := client.User.Query().Where(user.UsernameEQ("xxcheng123")).First(context.Background())
	if err != nil {
		panic(err)
	}
	cars, _ := u.QueryCars().All(context.Background())
	for _, car := range cars {
		fmt.Println(car)
	}
}
