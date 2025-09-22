package routes

import (
	"task4/controller"

	"github.com/gin-gonic/gin"
)

func LoginRoutesInit(router *gin.Engine) {
	loginRoutes := router.Group("/")
	{
		loginRoutes.POST("/login", controller.LoginController{}.Login)
	}

}
