package common

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/3 16:13
 */

type User struct {
	Username string
	Age      int
	Friends  []*User
}
