package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID      uint   `gorm:"column:userID" json:"userID"`
	UserName    string `gorm:"column:userName" json:"userName"`
	PhoneNumber string `gorm:"column:phone_no" json:"phone_no"`
	Password    uint   `gorm:"column:password" json:"password"`
	Email       string `gorm:"column:email" json:"email"`
}
