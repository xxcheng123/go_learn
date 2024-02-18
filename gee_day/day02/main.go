package main

import (
	"fmt"
	"gee"
	"net/http"
	"strconv"
	"time"
)

func main() {
	engine := gee.New()
	engine.GET("/time", func(c *gee.Context) {
		t := time.Now().Unix()
		c.String(http.StatusOK, strconv.FormatInt(t, 10))
	})
	engine.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello")
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
