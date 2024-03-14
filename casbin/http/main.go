package main

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-14 13:50:27
 * @LastEditTime: 2024-03-14 14:22:21
 */

func main() {
	e, err := casbin.NewEnforcer("../model.conf", "../policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	um := NewUm()
	um.AddUser(&User{
		Name: "admin",
		Role: "super::admin",
	}).AddUser(&User{
		Name: "user",
		Role: "user::common",
	}).AddUser(&User{
		Name: "xxcheng",
		Role: "user::common",
	})
	{
		api := r.Group("/api", func(ctx *gin.Context) {
			path := ctx.Request.URL.Path
			method := ctx.Request.Method
			userName := ctx.Request.Header.Get("user")
			user, err := um.FindUserByName(userName)
			ctx.Set("user", user)
			if err != nil {
				ctx.AbortWithStatus(401)
				return
			}
			if ok, err := e.Enforce(user.Role, path, method); !ok || err != nil {
				ctx.AbortWithStatus(403)
				return
			}
		})
		api.GET("/user/info", func(ctx *gin.Context) {
			u, _ := ctx.Get("user")
			user := u.(*User)
			ctx.JSON(200, gin.H{
				"msg":  "success",
				"name": user.Name,
				"role": user.Role,
			})
		})
		api.GET("/admin/info", func(ctx *gin.Context) {
			u, _ := ctx.Get("user")
			user := u.(*User)
			ctx.JSON(200, gin.H{
				"msg":  "success",
				"name": user.Name,
				"role": user.Role,
			})
		})
	}
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type UserManager interface {
	FindUserByName(name string) (*User, error)
	AddUser(*User) UserManager
}

type um struct {
	list map[string]*User
}

func (um *um) FindUserByName(name string) (*User, error) {
	if user, ok := um.list[name]; ok {
		return user, nil
	}
	return nil, errors.New("not found")
}
func (um *um) AddUser(user *User) UserManager {
	if _, ok := um.list[user.Name]; ok {
		return um
	}
	um.list[user.Name] = user
	return um
}

func NewUm() UserManager {
	return &um{
		list: make(map[string]*User),
	}
}
