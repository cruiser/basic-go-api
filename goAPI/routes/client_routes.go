package routes

import (
	"github.com/gin-gonic/gin"
	"simplyChoreo_Backend/handlers"
	"simplyChoreo_Backend/middleware"
)

func RegisterClientRoutes(r *gin.Engine) {
	clients := r.Group("/api/clients")
	clients.Use(middleware.AuthMiddleware())
	{
		clients.GET("/", handlers.GetAllClients)
		clients.POST("/", handlers.CreateClient)
	}
}
