package main

import "C"
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 中间件
	//中间件分为：全局中间件、分组中间件、单路由中间件
	// 运行时把中间件放到最前面，否则有些将无法注册上
	r.Use(func(c *gin.Context) {
		fmt.Printf("the reqeust url:%s\n", c.Request.URL.Path)
	})
	// 基础
	// 方法 POST / GET / PUT / DELETE / PATCH / HEAD / OPTIONS
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 上下文
	r.Any("/hello", func(c *gin.Context) {
		// http 报文 请求行/请求头/空行/请求数据
		// 读/写 header
		userAgent := c.Request.Header.Get("User-Agent")
		contentType := c.Request.Header.Get("Content-Type")
		// url query
		user := c.Request.URL.Query().Get("user")
		// 读 form query param
		//先解析
		// 请求有不同的结果，比如
		// form-data、x-www-form-urlencoded、json
		//对于json，需要使用BindJson来解析
		err := c.Request.ParseForm()
		if err != nil {
			log.Print(err)
		}
		token := c.Request.PostForm.Get("token")
		//优先级问题
		//如果url和body，有相同的key，如何处理？
		// 返回数据
		//返回的数据，包括 header 和 body
		//中间件的上下文传递
		c.Set("role", "super_admin")
		c.Header("X-Response-Time", strconv.FormatInt(time.Now().UnixMilli(), 10))
		c.JSON(200, gin.H{
			"message":      "hello",
			"timestamp":    time.Now().UnixMilli(),
			"user-agent":   userAgent,
			"user":         user,
			"token":        token,
			"content-type": contentType,
		})
	})
	// 路由
	//路由主要的核心功能就是，注册和查找
	// 注册、前缀树
	r.GET("/user/:id", func(c *gin.Context) {
		uid := c.Param("id")
		if uid != "" {
			c.String(http.StatusOK, fmt.Sprintf("Hello %s", uid))
		} else {
			//这个else是不可能执行到的
			c.String(http.StatusOK, fmt.Sprintf("Please input your user id."))
		}
	})
	// 分组
	v1 := r.Group("/v1")
	v1.Use(func(c *gin.Context) {
		fmt.Printf("[v1]the reqeust url:%s\n", c.Request.URL.Path)
	})
	{
		v1.GET("/a", func(c *gin.Context) {
			c.String(http.StatusOK, "v1.a")
		})
		v1.GET("/b", func(c *gin.Context) {
			c.String(http.StatusOK, "v1.b")
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/a", func(c *gin.Context) {
			c.String(http.StatusOK, "v2.a")
		})
		v2.GET("/b", func(c *gin.Context) {
			c.String(http.StatusOK, "v2.b")
		})
	}

	r.GET("md", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}, func(c *gin.Context) {
		fmt.Printf("[single]the reqeust url:%s\n", c.Request.URL.Path)
	})
	// 错误处理
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
