package controllers

import (
	"flutter_login_api_golang/config"
	"flutter_login_api_golang/models"
	"net/http"
	"os"
	"time"

	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Login(c *gin.Context) {
	var request struct {
		NomorInduk string `json:"nomor_induk"`
		Password   string `json:"password"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Permintaan tidak valid",
		})
		return
	}

	var siswa models.Siswa
	result := config.DB.Where("nomor_induk = ?", request.NomorInduk).First(&siswa)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Login gagal, periksa nomor induk dan password",
		})
		return
	}

	hash := md5.Sum([]byte(request.Password))
	hashedInput := hex.EncodeToString(hash[:])

	if hashedInput != siswa.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Login gagal, periksa nomor induk dan password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"siswa_id": siswa.SiswaID,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal membuat token",
		})
		return
	}

	siswa.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login berhasil",
		"data": gin.H{
			"token": tokenString,
			"siswa": siswa,
		},
	})
}
