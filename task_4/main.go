package main

import (
	"task4/core"
	"task4/middleware"
	"task4/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	core.Doconnect()

}

func main() {
	r := gin.Default()
	core.InitFile("logrus_study/gin_logrus/logs", "server")
	r.Use(middleware.LogMiddleware())
	routes.LoginRoutesInit(r)
	routes.RegisterRoutesInit(r)
	routes.PostRoutesInit(r)
	r.Run(":8080")
}
