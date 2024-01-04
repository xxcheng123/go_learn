package main

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/4 11:06
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	FirstName   string `json:"first_name" column:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func main() {
	db, err := sql.Open("mysql", "root:57MBLYOs2joES1bG@tcp(localhost:3306)/atguigudb")
	defer func() {
		_ = db.Close()
	}()
	fmt.Println(db, err)
	rows, err := db.Query("select `employee_id`,`first_name` from `employees`;")
	fmt.Println(err)
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		fmt.Println(err, id, firstName)
	}

}
