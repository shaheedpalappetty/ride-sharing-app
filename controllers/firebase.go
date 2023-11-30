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

// Diasble Status Notifications
func DisablePushNotifications(c *gin.Context) {
	var firebase models.Firebase
	if err := c.Bind(&firebase); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Bind Data",
		})
		return
	}
	if err := database.DB.Model(&models.Firebase{}).Where("user_id = ? AND category = ?", firebase.UserId, firebase.Category).Updates(map[string]interface{}{
		"token":  firebase.Token,
		"status": "InActive"}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update data",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Data updated successfully",
	})
}
