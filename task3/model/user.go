package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        int
	Name          string
	TotalPostsNum int
	Post          []Post `gorm:"foreignKey:UserID;references:UserID"`
}

func (User) TableName() string {
	return "users"
}
