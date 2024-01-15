package main

import "database/sql"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/11 9:36
 */

func main() {
	db, err := sql.Open("mysql", "dsn")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
