package controllers

import (
	"strconv"
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

func EditUserDetails(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Bind Data",
		})
	}
	data := map[string]interface{}{
		"full_name":    user.FullName,
		"birth_date":   user.BirthDate,
		"email":        user.Email,
		"password":     user.Password,
		"phone_number": user.PhoneNumber,
		"gender":       user.Gender,
		"image":        user.Image,
	}
	if err := database.DB.Model(&models.User{}).Where("id=?", user.ID).Updates(data).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Update in Database",
		})
	}
	c.JSON(200, gin.H{
		"success": "Updated Succesfully",
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Where("id=?", id).Delete(&models.User{}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Delete User",
		})
	}
	c.JSON(200, gin.H{
		"success": "Deleted User Details",
	})
}

func Payment(c *gin.Context) {
	var input struct {
		UserId   int    `json:"userid"`
		DriverId int    `json:"driverid"`
		Date     string `json:"date"`
		Status   string `json:"status"`
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
	if err := database.DB.Model(&models.Payment{}).Where("booking_id", booking.ID).Update("status", input.Status).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to Upadte Payment Deatails",
		})
	}
	c.JSON(200, gin.H{
		"message": "successfully added payment details",
	})
}

type BookingDetails struct {
	ID        int
	BookingID int
	User_id   int
	Date      string
	DriverId  int
	Distance  float64
	Fare      int
	Status    string
	Name      string
	DriverImg string
}

func CompletedTrips(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))

	var payments []BookingDetails
	if err := database.DB.Table("payments").Select("payments.id,payments.booking_id,payments.user_id,payments.date,payments.driver_id,payments.distance,payments.fare,payments.status,drivers.name,drivers.driver_img").
		Joins("INNER JOIN drivers ON drivers.id=payments.driver_id").Where("payments.user_id=? AND payments.status=?", id, "Paid").Scan(&payments).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	c.JSON(200, gin.H{
		"payments": payments,
	})

}

func ActiveTrips(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))
	var payments []BookingDetails
	if err := database.DB.Table("payments").Select("payments.id,payments.booking_id,payments.user_id,payments.date,payments.driver_id,payments.distance,payments.fare,payments.status,drivers.name,drivers.driver_img").
		Joins("INNER JOIN drivers ON drivers.id=payments.driver_id").Where("payments.user_id=? AND payments.status=?", id, "pending").Scan(&payments).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	c.JSON(200, gin.H{
		"payments": payments,
	})
}
