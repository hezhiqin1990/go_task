package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	UserID   uint
	Comments []Comment `gorm:"foreignkey:PostID;references:ID"`
}

func (Post) TableName() string {
	return "post"
}
