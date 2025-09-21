package main

import (
	"fmt"
	"task3/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:6235135@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	//DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	//addPost()
	//addComments()
	//funddes()

	//maxComment()
	//PostAdd()
	CommentDelete()
}

func adduser() {
	user := model.User{
		UserID: 2,
		Name:   "Keen",
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	result := DB.Create(&user) //通过数据的指针来创建
	if result.RowsAffected > 1 {
		fmt.Print(user.Model.ID)
	}
	fmt.Println(result.RowsAffected)
	fmt.Println(user.UserID)

}

func addPost() {
	post := model.Post{
		UserID:  2,
		Title:   "啊实打实大师奥术大师大所大所多asdasdasdasdasd",
		Content: "啊实打实大师奥术大师大所大所多asd asdasdasd",
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	result := DB.Create(&post) //通过数据的指针来创建
	if result.RowsAffected > 1 {
		fmt.Print(post.UserID)
	}
	fmt.Println(result.RowsAffected)
	fmt.Println(post.UserID)

}

func addComments() {
	post := model.Comment{
		UserID:  2,
		Content: "as打算的撒你大爷asdasd",
		PostID:  2,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	result := DB.Create(&post) //通过数据的指针来创建
	if result.RowsAffected > 1 {
		fmt.Print(post.UserID)
	}
	fmt.Println(result.RowsAffected)
	fmt.Println(post.UserID)

}

/*
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息
*/
func funddes() {

	var users []model.User
	DB.Preload("Post").Preload("Post.Comment").Where("user_id=?", 2).Find(&users)
	fmt.Println(users)

}

/*
	使用Gorm查询评论数量最多的文章信息
*/

func maxComment() {

	var post model.Post
	DB.Order("total_comments_num Desc").Limit(1).Find(&post)
	fmt.Println(post)

}

//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。

func PostAdd() {

	post := model.Post{
		UserID:  2,
		Title:   "啊实打实大师奥术大师大所大所多asdasdasdasdasd",
		Content: "啊实打实大师奥术大师大所大所多asd asdasdasd",
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	result := DB.Create(&post) //通过数据的指针来创建
	fmt.Println(result.RowsAffected)
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func CommentDelete() {
	comment := model.Comment{
		PostID: 7,
		Model: gorm.Model{
			ID: 1,
		},
	}
	result := DB.Delete(&comment) //通过数据的指针来创建
	fmt.Println(result.RowsAffected)
}
