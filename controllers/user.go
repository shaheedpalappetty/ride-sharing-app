package controllers

import (
	"taxi_app/database"
	"taxi_app/models"

	"github.com/gin-gonic/gin"
)

// Create User
func SignUpUser(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	if err := validate.Struct(user); err != nil {
		c.JSON(401, gin.H{
			"error": err,
		})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add detailes in database",
		})
		return
	}
	c.JSON(200, user.ID)
}

// Login User(Sign In)
func LoginUser(c *gin.Context) { 
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to bind",
		})
		return
	}
	var user models.User
	if err := database.DB.Where("email =?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find user",
		})
		return
	}
	if user.Email != credentials.Email || user.Password != credentials.Password {
		c.JSON(400, gin.H{
			"error": "incorrect username or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetDriversListAccordingtoVehicleType(c *gin.Context){
	
}

