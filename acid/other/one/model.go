package main

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/22 15:14
 */

type User struct {
	Id       int    `gorm:"id"`
	Username string `gorm:"username"`
	Grade    int    `gorm:"grade"`
}
