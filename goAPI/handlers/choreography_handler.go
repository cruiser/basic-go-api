package handlers

import (
	"net/http"
	"simplyChoreo_Backend/database"
	"simplyChoreo_Backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllChoreography(c *gin.Context) {
	var choreography []models.Choreography
	result := database.DB.Find(&choreography)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No choreography found"})
	}

	c.JSON(http.StatusOK, choreography)
}

func CreateChoreography(c *gin.Context) {
	var choreography models.Choreography

	if err := c.ShouldBindJSON(&choreography); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&choreography)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, choreography)
}
