package controllers

import (
	"log"
	"math"
	"sort"
	"strconv"
	"taxi_app/database"
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
	nearestDrivers, err := findNearestDrivers(booking.PickupLat, booking.PickupLong, 3)
	//get the driver details
	var driver models.Detail

	var drivers []models.Detail

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

type Location struct {
	ID        int
	Latitude  float64
	Longitude float64
}

func findNearestDrivers(PickupLat float64, PickupLong float64, count int) ([]Location, error) {
	distanceMap := make(map[float64]Location)
	// get driver Locations
	var driverLocations []Location
	if err := database.DB.Table("drivers").Select("id", "latitude", "longitude").Scan(&driverLocations).Error; err != nil {
		return nil, err
	}
	log.Println(driverLocations)
	// Calculate distances for each driver
	for _, driverLocation := range driverLocations {
		distance := haversine(PickupLat, PickupLat, driverLocation.Latitude, driverLocation.Longitude)
		distanceMap[distance] = driverLocation
	}

	// Sort distances
	distances := make([]float64, 0, len(distanceMap))
	for distance := range distanceMap {
		distances = append(distances, distance)
	}
	sort.Float64s(distances)

	// Select nearest drivers
	nearestDrivers := make([]Location, 0, count)
	for i := 0; i < count && i < len(distances); i++ {
		nearestDrivers = append(nearestDrivers, distanceMap[distances[i]])
	}

	return nearestDrivers, nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // Earth radius in kilometers

	// Convert degrees to radians
	lat1Rad := toRadians(lat1)
	lon1Rad := toRadians(lon1)
	lat2Rad := toRadians(lat2)
	lon2Rad := toRadians(lon2)

	// Calculate differences
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := earthRadius * c

	return distance
}

func toRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
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
