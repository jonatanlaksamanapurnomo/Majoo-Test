package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(45);" json:"name"`
	UserName  string `gorm:"type:varchar(45)" json:"user_name"`
	Password  string `gorm:"size:128" json:"password"`
	UpdatedBy int    `gorm:"type:int" json:"updated_by"`
	CreatedBy int    `gorm:"type:int" json:"created_by"`
}

func GetUser() User {
	var user User
	return user
}

func GetUsers() []User {
	var users []User
	return users
}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
