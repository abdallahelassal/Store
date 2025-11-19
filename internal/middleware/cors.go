package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc{
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET","POST","DELETE","PUT","PATCH","OPTIONS"}
	config.AllowHeaders = []string{"Origin","Content-Type","Accept","Authorization"}
	config.AllowCredentials = false
	return cors.New(config)
}