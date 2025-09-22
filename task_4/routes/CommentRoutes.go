package routes

import (
	"task4/controller"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRoutesInit(router *gin.Engine) {
	commentRoutes := router.Group("/")
	{
		commentRoutes.POST("/comment", middleware.JwtAuthMiddleware(), controller.CommentController{}.Add)
	}

}
