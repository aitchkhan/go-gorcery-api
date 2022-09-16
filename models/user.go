package models

import (
	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `gorm:"type:varchar(200);uniqueIndex"`
	HashedPassword string `json:"-"`
	Address string `json:"address"`
}

func CreateUser(db *gorm.DB, u *User) {
	db.Create(u)
}

func GetUserByEmail(db *gorm.DB, u *User) {
	spew.Dump("in repo", u)
	db.Where("email = ?", u.Email).First(&u)
}