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

// Check Driver status by Mobile Number
func CheckDriverStatus(c *gin.Context) {
	number := c.Param("number")

	var driver models.Driver
	if err := database.DB.Where("phone_number=?", number).First(&driver).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find user",
		})
		return
	}
	c.JSON(200, driver.Status)
}

// Update Location of driver
func UpdateLocation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var position struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}
	if err := c.Bind(&position); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to Bind data",
		})
		return
	}

	body := map[string]interface{}{
		"latitude":  position.Latitude,
		"longitude": position.Longitude,
	}

	if err := database.DB.Model(&models.Driver{}).Where("id = ?", id).Updates(body).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to update details in database",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Updated Location Successfully",
	})

}

func StartTrip(c *gin.Context) {
	var input struct {
		SLatitude  float64 `json:"slatitude"`
		SLongitude float64 `json:"slongitude"`
		UserId     int     `json:"userid"`
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Bind data",
		})
		return
	}

	var booking models.Booking
	if err := database.DB.Where("user_id = ?", input.UserId).Last(&booking).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to fetch data from database",
		})
		return
	}

	if err := database.DB.Model(&models.Booking{}).Where("id=?", booking.ID).Updates(map[string]interface{}{
		"pickup_lat":  input.SLatitude,
		"pickup_long": input.SLongitude,
	}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Update Location in Bookings Table",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Updated pickUp Location in booking table",
	})
}

func Revenue(c *gin.Context) {
	var input struct {
		Driver_id int    `json:"driverid"`
		Date      string `json:"date"`
	}
	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get data",
		})
		return
	}
	var payments []models.Payment
	if err := database.DB.Where("driver_id = ? AND date = ? AND status=?", input.Driver_id, input.Date, "Paid").Find(&payments).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to find data",
		})
		return
	}
	c.JSON(200, gin.H{
		"payments": payments,
	})

}

func ChangeOnlineStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	status := c.Query("status")

	var accepted bool
	if status == "t" {
		accepted = true
	} else {
		accepted = false
	}

	if err := database.DB.Model(&models.Driver{}).Where("id=?", id).Update("online", accepted).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update online status",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "status updated",
	})
}