package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func OpenDBConn() {
	dbString := "root:root@tcp(127.0.0.1:3306)/go_grocery_local?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
}