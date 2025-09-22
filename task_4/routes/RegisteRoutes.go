package routes

import (
	"task4/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesInit(router *gin.Engine) {
	registerRoutes := router.Group("/")
	{
		registerRoutes.POST("/register", controller.RegisterController{}.Register)
	}

}
