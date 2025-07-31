package main

import (
	"flutter_login_api_golang/config"
	"flutter_login_api_golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	r.POST("/api/login", controllers.Login)

	r.Run(":8080") // http://localhost:8080
}
