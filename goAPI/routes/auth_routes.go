package routes

import (
	"github.com/gin-gonic/gin"
	"simplyChoreo_Backend/handlers"
)

func RegisterAuthRoutes(a *gin.Engine) {
	auth := a.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}
}
