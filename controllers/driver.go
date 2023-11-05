package controllers

import (
	"strconv"
	"taxi_app/database"
	"taxi_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

// add Driver
func AddDriver(c *gin.Context) {
	var driver models.Driver
	if err := c.Bind(&driver); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	if err := validate.Struct(driver); err != nil {
		c.JSON(401, gin.H{
			"error": err,
		})
		return
	}
	driver.Status = "Pending"
	if err := database.DB.Create(&driver).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add detailes in database",
		})
		return
	}
	c.JSON(200, driver.ID)
}

// Login Driver
func Driverlogin(c *gin.Context) {

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to bind",
		})
		return
	}
	var driver models.Driver
	if err := database.DB.Where("user_name =?", credentials.Username).First(&driver).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find user",
		})
		return
	}
	if driver.UserName != credentials.Username || driver.Password != credentials.Password {
		c.JSON(400, gin.H{
			"error": "incorrect username or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"driver": driver,
	})

}

// Add User Documents
func AddDocuments(c *gin.Context) {
	var documents models.DriverDocuments
	if err := c.Bind(&documents); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}

	if err := database.DB.Create(&documents).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add detailes in database",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "successfully added driver Documents",
	})
}

// Add Vehicle Details
func AddVehicleDetails(c *gin.Context) {

	var VehicleDetails models.VehicleDetails
	if err := c.Bind(&VehicleDetails); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}

	if err := database.DB.Create(&VehicleDetails).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add detailes in database",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "successfully added Vehicle Details",
	})
}

// Get Driver Details
func GetDriverDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("driver_id"))
	var driver models.Driver
	if err := database.DB.First(&driver, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find user",
		})
		return
	}
	c.JSON(200, gin.H{
		"driver": driver,
	})
}
