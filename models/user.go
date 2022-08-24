package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string
}


func CreateUser(db *gorm.DB, u *User) {
	db.Create(u)
}

func Signin(db *gorm.DB, u *User) {
	db.First(u, u.ID)
}