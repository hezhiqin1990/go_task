package controller

import (
	"net/http"
	"task4/core"
	"task4/model"
	"time"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	BaseController
}

func (p PostController) Add(ctx *gin.Context) {
	var post model.Post
	if err := ctx.ShouldBindJSON(&post); err == nil {
		if len(post.Title) == 0 || len(post.Content) == 0 {
			p.Error(ctx, gin.H{
				"code": 400,
				"msg":  "文章标题，或内容不能为空",
			})
			return
		} else {
			core.DB.Create(&post)
			p.Success(ctx, gin.H{
				"code": http.StatusOK,
				"msg":  "文章保存成功",
			})

		}

	} else {
		p.Error(ctx, gin.H{
			"code": 400,
			"msg":  err,
		})
	}
}

func (p PostController) Get(ctx *gin.Context) {
	postId := ctx.Query("post_id")
	post := model.Post{}
	core.DB.Preload("Comments").Where("id=?", postId).Find(&post)
	p.Success(ctx, post)
}

func (p PostController) Update(ctx *gin.Context) {
	var post model.Post
	var selectPost model.Post
	if err := ctx.ShouldBindJSON(&post); err == nil {
		if post.ID == 0 {
			p.Error(ctx, gin.H{
				"code": 400,
				"msg":  "文章id不存在，请重试",
			})
			return
		}
		core.DB.Where("id=?", post.ID).Find(&selectPost)
		if selectPost.UserID != post.UserID {
			p.Error(ctx, gin.H{
				"code": 400,
				"msg":  "仅允许文章的作者修改，请重试",
			})
			return
		} else {
			if len(post.Title) == 0 || len(post.Content) == 0 {
				p.Error(ctx, gin.H{
					"code": 400,
					"msg":  "文章标题，或内容不能为空",
				})
				return
			} else {
				post.CreatedAt = time.Now()
				core.DB.Save(&post)
				p.Success(ctx, gin.H{
					"code": http.StatusOK,
					"msg":  "文章修改成功",
				})

			}
		}

	} else {
		p.Error(ctx, gin.H{
			"code": 400,
			"msg":  err,
		})
	}
}

func (p PostController) Delete(ctx *gin.Context) {
	postId := ctx.Query("post_id")
	userid, _ := ctx.Get("userid")
	var selectPost model.Post
	core.DB.Where("id = ?", postId).Find(&selectPost)
	if selectPost.UserID != userid {
		p.Error(ctx, gin.H{
			"code": 400,
			"msg":  "仅允许文章的作者删除，请重试",
		})
		return
	}
	core.DB.Delete(&model.Post{}, postId)
	p.Success(ctx, gin.H{
		"code": http.StatusOK,
		"msg":  "文章删除成功",
	})
}
