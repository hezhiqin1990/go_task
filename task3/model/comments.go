package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	User    User `gorm:"foreignKey:UserID;references:UserID"`
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID;references:ID"`
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	post := Post{}
	result := tx.Model(&Post{}).Where("id = ?", c.PostID).
		UpdateColumn("total_comments_num", gorm.Expr("total_comments_num -?", 1)).
		Scan(&post)
	if post.TotalCommentsNum <= 0 {
		post.Status = "无评论"
		tx.Save(&post)
	}
	return result.Error
}
