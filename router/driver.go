package router

import (
	"taxi_app/controllers"

	"github.com/gin-gonic/gin"
)

func DriverRouter(r *gin.Engine) {
	r.POST("/adddriver", controllers.AddDriver)
	r.POST("/adddocuments", controllers.AddDocuments)
	r.POST("/addvehicledetails", controllers.AddVehicleDetails)
	r.GET("/driverdetail", controllers.GetDriverDetail)

	
}
