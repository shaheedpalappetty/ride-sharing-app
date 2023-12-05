package controllers

import (
	"math"
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

func GetFare(c *gin.Context) {
	var values Coordinate
	if err := c.Bind(&values); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed To Bind Data",
		})
		return
	}
	fare := CalculateFare(values)
	c.JSON(200, gin.H{
		"success": fare,
	})

}

const (
	baseFare         = 50.0 // Base fare
	farePerKilometer = 13.5 // Fare per kilometer
)

// Coordinate represents a geographical coordinate with latitude and longitude.
type Coordinate struct {
	SLatitude  float64 `json:"slatitude"`
	SLongitude float64 `json:"slongitude"`
	ELatitude  float64 `json:"elatitude"`
	ELongitude float64 `json:"elongitude"`
}

// CalculateDistance calculates the distance between two coordinates using Haversine formula.
func CalculateDistance(coord Coordinate) float64 {
	const earthRadius = 6371 // Earth radius in kilometers

	// Convert latitude and longitude from degrees to radians
	lat1 := degToRad(coord.SLatitude)
	lon1 := degToRad(coord.SLongitude)
	lat2 := degToRad(coord.ELatitude)
	lon2 := degToRad(coord.ELongitude)

	// Calculate differences in coordinates
	dlat := lat2 - lat1
	dlon := lon2 - lon1

	// Haversine formula
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := earthRadius * c

	return distance
}

// degToRad converts degrees to radians.
func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// CalculateFare calculates the fare based on the distance.
func CalculateFare(cordinates Coordinate) float64 {
	distance := CalculateDistance(cordinates)
	fare := baseFare + distance*farePerKilometer
	return fare
}
