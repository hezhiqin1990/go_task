package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (c BaseController) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (c BaseController) Error(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
