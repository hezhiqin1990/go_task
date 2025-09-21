package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title            string
	Content          string
	UserID           uint
	TotalCommentsNum int
	Status           string
	Comment          []Comment `gorm:"foreignKey:PostID;references:ID"`
}

func (Post) TableName() string {
	return "posts"
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	userres := User{}
	result := tx.Model(&User{}).Where("user_id = ?", post.UserID).
		UpdateColumn("total_posts_num", gorm.Expr("total_posts_num + ?", 1)).
		Scan(&userres)
	fmt.Println("TotalCommentNum:", userres.TotalPostsNum)
	return result.Error
}
