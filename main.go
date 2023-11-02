package main

import (
	authController "authGin/controllers"
	"authGin/models"

	"github.com/gin-gonic/gin"
)

func main()  {
	Router := gin.Default()
	models.ConnectDatabase()

	Router.POST("/api/v1/auth/register", authController.Register)
	Router.POST("/api/v1/auth/login", authController.Login)

	Router.Run()
}