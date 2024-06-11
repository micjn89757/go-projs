package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 处理跨域问题
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000"}, // 等同于允许所有域名 AllowAllOrigins: true
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
			AllowHeaders: []string{"*", "Authorization"},
			ExposeHeaders: []string{"Content-Length", "test/plain", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge: 12*time.Hour,
		})
}