package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	{
		v1 := engine.Group("/v1")
		v1.GET("/user/:username", func(c *gee.Context) {
			c.String(http.StatusOK, fmt.Sprintf("v1:hello,%s", c.Param("username")))
		})
	}
	{
		v2 := engine.Group("/v2")
		v2.GET("/user/:username", func(c *gee.Context) {
			c.String(http.StatusOK, fmt.Sprintf("v2:hello,%s", c.Param("username")))
		})
	}
	fmt.Println("running...")
	engine.Run(":8080")
}
