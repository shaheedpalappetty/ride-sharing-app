package controllers

import (
	"strconv"
	"taxi_app/database"
	"taxi_app/helper"
	"taxi_app/models"

	"github.com/gin-gonic/gin"
)

func RequestRide(c *gin.Context) {
	var booking models.Booking
	if err := c.Bind(&booking); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to Bind data",
		})
		return
	}
	if err := validate.Struct(booking); err != nil {
		c.JSON(401, gin.H{
			"error": "Null Values",
		})
		return
	}

	if err := database.DB.Create(&booking).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add detailes in database",
		})
		return
	}

	//find nearest drivers
	nearestDrivers, err := helper.FindNearestDrivers(booking.PickupLat, booking.PickupLong, 3)
	//get the driver details
	var driver models.AvailableDrivers

	var drivers []models.AvailableDrivers

	for i := range nearestDrivers {
		if err := database.DB.Table("drivers").
			Select("drivers.id, drivers.name, drivers.last_name, drivers.phone_number, drivers.email, drivers.birth_date, drivers.driver_img, drivers.gender, drivers.experience, vehicle_details.vehicle_brand, vehicle_details.vehicle_model, vehicle_details.vehicle_year, vehicle_details.vehicle_color,vehicle_details.vehicle_type, vehicle_details.vehicle_seat, vehicle_details.vehicle_number").
			Joins("INNER JOIN vehicle_details ON vehicle_details.user_id = drivers.id").
			Where("drivers.id = ?", nearestDrivers[i].ID).
			Scan(&driver).Error; err != nil {
			c.JSON(400, gin.H{
				"error": "error fetching data",
			})
			return
		}
		fare := helper.CalculateFareWithVehicleType(models.Coordinate{
			SLatitude:  booking.PickupLat,
			SLongitude: booking.PickupLong,
			ELatitude:  booking.DropoffLat,
			ELongitude: booking.DropoffLong,
		}, driver.VehicleType)
		driver.Fare = fare
		drivers = append(drivers, driver)
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": drivers,
	})

}



// Get Token of Driver and User
func ConfirmRide(c *gin.Context) {
	id := c.Param("driver_id")
	var firebase models.Firebase
	if err := database.DB.Where("user_id = ? AND category = ? AND status = ?", id, "Driver", "Active").First(&firebase).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to find Data",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": firebase.Token,
	})
}
func UpdateRideStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("userid"))
	var status struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&status); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Bind Data",
		})
		return
	}
	if err := database.DB.Model(&models.Booking{}).Where("user_id = ?", id).Update("status", status.Status).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Update Data",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Succefully Updated data",
	})
}

func GetRideFare(c *gin.Context) {
	var input struct {
		ELatitude  float64 `json:"elatitude"`
		ELongitude float64 `json:"elongitude"`
		UserId     int     `json:"userid"`
		DriverId   int     `json:"driverid"`
		Date       string  `json:"date"`
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
		"dropoff_lat":  input.ELatitude,
		"dropoff_long": input.ELongitude,
	}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Update Location in Bookings Table",
		})
		return
	}
	var driver models.VehicleDetails
	if err := database.DB.Where("user_id = ?", input.DriverId).First(&driver).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}

	coordinate := models.Coordinate{
		SLatitude:  booking.PickupLat,
		SLongitude: booking.PickupLong,
		ELatitude:  input.ELatitude,
		ELongitude: input.ELongitude,
	}

	fare := helper.CalculateFareWithVehicleType(coordinate, driver.VehicleType)

	if err := database.DB.Create(&models.Payment{
		BookingID: booking.ID,
		User_id:   input.UserId,
		Date:      input.Date,
		DriverId:  driver.UserID,
		Distance:  helper.CalculateDistance(coordinate),
		Fare:      fare,
		Status:    "pending",
	}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create data",
		})
		return
	}

	c.JSON(200, gin.H{
		"fare": helper.CalculateFareWithVehicleType(coordinate, driver.VehicleType),
	})

}
