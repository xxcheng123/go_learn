package main

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/8 13:34
 */

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	DepartmentId sql.NullInt32 `json:"department_id"`
	Salary       float64       `gorm:"salary"`
	Count        int           `gorm:"count"`
}

func main() {
	dsn := "root:57MBLYOs2joES1bG@tcp(localhost:3306)/atguigudb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var employees = &[]Employee{}
	db.Select("department_id", "sum(salary) salary", "count(*) count").
		Group("department_id").
		Find(employees)
	for _, dep := range *employees {
		if dep.DepartmentId.Valid {
			fmt.Printf("id:[%d],total persons:[%d],total salary:[%f]\n", dep.DepartmentId.Int32, dep.Count, dep.Salary)
		}

	}
}
