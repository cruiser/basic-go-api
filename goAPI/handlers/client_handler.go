package handlers

import (
	"net/http"
	"simplyChoreo_Backend/database"
	"simplyChoreo_Backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllClients(c *gin.Context) {
	var clients []models.Client
	result := database.DB.Find(&clients)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Clients Found"})
	}

	c.JSON(http.StatusOK, clients)
}

func CreateClient(c *gin.Context) {
	var client models.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&client)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}
