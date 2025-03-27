package routes

import (
	"github.com/gin-gonic/gin"
	"simplyChoreo_Backend/handlers"
	"simplyChoreo_Backend/middleware"
)

func RegisterChoreographyRoutes(r *gin.Engine) {
	choreo := r.Group("/api/choreography")
	choreo.Use(middleware.AuthMiddleware())
	{
		choreo.GET("/", handlers.GetAllChoreography)
		choreo.POST("/", handlers.CreateChoreography)
	}
}
