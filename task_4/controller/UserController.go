package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (c UserController) Index(ctx *gin.Context) {
	var user = map[string]string{
		"username": "IT营小王子",
		"password": "123456",
	}

	c.Success(ctx, user)
}
func (c UserController) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "增加用户")
}
