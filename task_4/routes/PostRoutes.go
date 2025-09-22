package routes

import (
	"task4/controller"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutesInit(router *gin.Engine) {
	postRoutes := router.Group("/post")
	{
		postRoutes.GET("/get", controller.PostController{}.Get)
		postRoutes.POST("/add", middleware.JwtAuthMiddleware(), controller.PostController{}.Add)
		postRoutes.POST("/update", middleware.JwtAuthMiddleware(), controller.PostController{}.Update)
		postRoutes.POST("/delete", middleware.JwtAuthMiddleware(), controller.PostController{}.Delete)
	}

}
