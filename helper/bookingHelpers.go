package helper

import (
	"log"
	"math"
	"sort"
	"taxi_app/database"
	"taxi_app/models"
)

// CalculateFareWithVehicleType calculates the fare based on the vehicle type.
func CalculateFareWithVehicleType(coordinates models.Coordinate, vehicleType string) int {
	// Update the base fare and fare per kilometer based on the vehicle type
	var baseFare float64
	var farePerKilometer float64

	switch vehicleType {
	case "Hatchback":
		baseFare = 50.0
		farePerKilometer = 10.0
	case "Sedan":
		baseFare = 60.0
		farePerKilometer = 12.0
	case "SUV":
		baseFare = 70.0
		farePerKilometer = 15.0
	default:
		// Use default values for unknown vehicle types
		baseFare = 50.0
		farePerKilometer = 10.0
	}

	// Calculate distance and fare
	distance := CalculateDistance(coordinates)
	fare := baseFare + distance*farePerKilometer
	rate := int(math.Round(fare))
	return rate
}

// CalculateDistance calculates the distance between two coordinates using Haversine formula.
func CalculateDistance(coord models.Coordinate) float64 {
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

func FindNearestDrivers(PickupLat float64, PickupLong float64, count int) ([]models.Location, error) {
	distanceMap := make(map[float64]models.Location)
	// get driver Locations
	var driverLocations []models.Location
	if err := database.DB.Table("drivers").Select("id", "latitude", "longitude").Scan(&driverLocations).Error; err != nil {
		return nil, err
	}
	log.Println(driverLocations)
	// Calculate distances for each driver
	for _, driverLocation := range driverLocations {
		distance := CalculateDistance(models.Coordinate{
			SLatitude:  PickupLat,
			SLongitude: PickupLong,
			ELatitude:  driverLocation.Latitude,
			ELongitude: driverLocation.Longitude,
		})
		distanceMap[distance] = driverLocation
	}

	// Sort distances
	distances := make([]float64, 0, len(distanceMap))
	for distance := range distanceMap {
		distances = append(distances, distance)
	}
	sort.Float64s(distances)

	// Select nearest drivers
	nearestDrivers := make([]models.Location, 0, count)
	for i := 0; i < count && i < len(distances); i++ {
		nearestDrivers = append(nearestDrivers, distanceMap[distances[i]])
	}

	return nearestDrivers, nil
}
