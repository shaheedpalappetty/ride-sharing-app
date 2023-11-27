package main

import (
	"log"
	"taxi_app/database"
	"taxi_app/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	database.ConnectToDB()
}

func main() {
	r := gin.Default()
	router.DriverRouter(r)
	router.AdminRouter(r)
	router.UserRouter(r)
	router.BookingRouter(r)
	router.FirebaseRouter(r)
	r.Run(":8080")
}
