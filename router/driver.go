package router

import (
	"taxi_app/controllers"

	"github.com/gin-gonic/gin"
)

func DriverRouter(router *gin.Engine) {
	r := router.Group("/driver")
	r.POST("/create", controllers.AddDriver)
	r.POST("/documents", controllers.AddDocuments)
	r.POST("/vehicledetails", controllers.AddVehicleDetails)
	r.GET("/detail", controllers.GetDriverDetail)

}
