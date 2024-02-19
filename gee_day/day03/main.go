package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/static/*file", func(c *gee.Context) {
		filename := c.Param("file")
		c.String(http.StatusOK, filename)
	})
	engine.GET("/user/:username", func(c *gee.Context) {
		c.String(http.StatusOK, fmt.Sprintf("hello,%s", c.Param("username")))
	})
	engine.GET("/", func(c *gee.Context) {
		c.JSON(http.StatusOK, struct {
			Username string
			Password string
		}{
			Username: "admin",
			Password: "123456",
		})
	})
	fmt.Println("running...")
	engine.Run(":8080")
}
