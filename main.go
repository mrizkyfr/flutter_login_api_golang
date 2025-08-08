package main

import (
	"flutter_login_api_golang/config"
	"flutter_login_api_golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	// Tambahkan CORS middleware
	r.Use(CORSMiddleware())

	r.POST("/api/login", controllers.Login)
	r.POST("/presensi/siswa", controllers.SiswaPresensi)
	r.POST("/presensi/do_presensi_siswa", controllers.DoPresensiSiswa)

	// r.Run(":8080") // http://localhost:8080
	r.Run()
}
