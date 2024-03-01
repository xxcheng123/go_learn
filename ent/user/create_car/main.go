package main

import (
	"context"
	"ent-demo/db"
	"fmt"
	"time"
)

func main() {
	client := db.Get()
	c1, err1 := client.Car.Create().SetID(time.Now().UnixMicro()).SetName("Tesla-1").SetPrice(99.99).SetCreateAt(time.Now()).Save(context.Background())
	c2, err2 := client.Car.Create().SetID(time.Now().UnixMicro()).SetName("BMW-2").SetPrice(919.99).SetCreateAt(time.Now()).Save(context.Background())
	fmt.Println(c1, err1, c2, err2)
}
