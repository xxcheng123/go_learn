package main

import (
	"fmt"
	"gee"
	"net/http"
	"time"
)

func main() {
	engine := gee.New()
	engine.Use(func(ctx *gee.Context) {
		start := time.Now()
		fmt.Println("global router has request....")
		ctx.Next()
		fmt.Println("cost time:", time.Since(start))
	}, gee.Recovery())
	{
		v1 := engine.Group("/v1")
		v1.Use(func(ctx *gee.Context) {
			fmt.Println("v1 router has request....")
		})
		v1.GET("/user/:username", func(c *gee.Context) {
			c.String(http.StatusOK, fmt.Sprintf("v1:hello,%s", c.Param("username")))
		})
	}
	{
		v2 := engine.Group("/v2")
		v2.Use(func(ctx *gee.Context) {
			fmt.Println("v1 router has request....")
		})
		v2.GET("/user/:username", func(c *gee.Context) {
			c.String(http.StatusOK, fmt.Sprintf("v2:hello,%s", c.Param("username")))
		})
	}
	engine.GET("/panic", func(ctx *gee.Context) {
		panic("eeee")
	})
	fmt.Println("running...")
	engine.Run(":8080")
}
