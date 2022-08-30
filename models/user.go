package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(db *gorm.DB, u *User) {
	db.Create(u)
}

func Signin(db *gorm.DB, u *User) {
	db.First(u, u.ID)
}