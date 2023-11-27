package controllers

import (
	"taxi_app/database"
	"taxi_app/models"

	"github.com/gin-gonic/gin"
)

func FirebaseCredentials(c *gin.Context) {
	var firebase models.Firebase
	if err := c.Bind(&firebase); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Bind Data",
		})
		return
	}
	if err := database.DB.Create(&firebase).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to Add Details in Database",
		})
		return
	}
	c.JSON(200, gin.H{
		"Success": "Successfully Created Firebase Model",
	})
}


//Update Status
func UpdateStatus(c *gin.Context){
	// var status string
}
