package router

import (
	"taxi_app/controllers"

	"github.com/gin-gonic/gin"
)

func DriverRouter(router *gin.Engine) {
	r := router.Group("/driver")
	{
		r.POST("/create", controllers.AddDriver)
		r.POST("/documents", controllers.AddDocuments)
		r.POST("/vehicledetails", controllers.AddVehicleDetails)
		r.GET("/detail", controllers.GetDriverDetail)
		r.POST("/login", controllers.Driverlogin)
		r.POST("/status/:number", controllers.CheckDriverStatus)
		r.PATCH("/updatelocation/:id", controllers.UpdateLocation) //change to ws
		r.PATCH("/starttrip", controllers.StartTrip)
		r.POST("/revenue", controllers.Revenue)
		r.PATCH("/online", controllers.ChangeOnlineStatus)
	}

}
