package controllers

import (
	"errors"
	"strconv"
	"taxi_app/database"
	"taxi_app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
func UpdateToken(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var body struct {
        Token string `json:"token"`
        User  string `json:"user"`
    }
    if err := c.Bind(&body); err != nil {
        c.JSON(400, gin.H{
            "error": "Failed to Bind Data",
        })
        return
    }

    var firebaseModel models.Firebase
    if err := database.DB.Where("user_id = ? AND category = ?", id, body.User).First(&firebaseModel).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(404, gin.H{
                "error": "User not found in the database with the given category",
            })
            return
        }
        c.JSON(500, gin.H{
            "error": "Failed to query database",
        })
        return
    }

    if err := database.DB.Model(&models.Firebase{}).Where("user_id = ? AND category = ?", id, body.User).Update("token", body.Token).Error; err != nil {
        c.JSON(500, gin.H{
            "error": "Failed to update details in the database",
        })
        return
    }

    c.JSON(200, gin.H{
        "success": "Firebase Token updated successfully",
    })
}
