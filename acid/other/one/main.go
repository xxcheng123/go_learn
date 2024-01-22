package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/22 15:05
 */

func main() {
	dsn := "gorm:12345678@tcp(localhost:3306)/tim?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	u := &User{
		Id:       1,
		Username: "22222",
		//Grade:    33,
	}
	tx := db.Table("user").Begin()
	result := tx.Save(u)
	fmt.Println(result.Error, result.RowsAffected)
	fmt.Println("休眠3s")
	time.Sleep(7 * time.Second)
	tx.Commit()
	fmt.Println("success.")
}
