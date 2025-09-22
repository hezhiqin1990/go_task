package controller

import (
	"net/http"
	"task4/core"
	"task4/model"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	BaseController
}

func (p CommentController) Add(ctx *gin.Context) {
	var c model.Comment
	if err := ctx.ShouldBindJSON(&c); err == nil {
		if len(c.Content) == 0 {
			p.Error(ctx, gin.H{
				"code": 400,
				"msg":  "内容不能为空",
			})
			return
		} else {
			userid := ctx.GetUint("userid")
			c.UserID = (userid)
			core.DB.Create(&c)
			p.Success(ctx, gin.H{
				"code": http.StatusOK,
				"msg":  "评论保存成功",
			})

		}

	} else {
		p.Error(ctx, gin.H{
			"code": 400,
			"msg":  err,
		})
	}
}
