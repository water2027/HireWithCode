package middleware

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterMiddleware(router *gin.Engine) {
	origin := os.Getenv("FRONTEND_URL")
	if origin == "" {
		origin = "http://localhost:5173"
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{origin},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
}
