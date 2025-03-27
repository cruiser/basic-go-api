package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"simplyChoreo_Backend/database"
	"simplyChoreo_Backend/routes"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://simplychoreo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	database.InitDB()

	routes.RegisterClientRoutes(r)
	routes.RegisterAuthRoutes(r)
	routes.RegisterChoreographyRoutes(r)

	err := r.Run(":8081")
	if err != nil {
		return
	}
}
