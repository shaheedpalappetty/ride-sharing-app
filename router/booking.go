package router

import (
	"taxi_app/controllers"

	"github.com/gin-gonic/gin"
)

func BookingRouter(router *gin.Engine) {
	r := router.Group("/booking")
	{
		r.POST("/ride", controllers.RequestRide)
		r.GET("/confirm/:driver_id", controllers.ConfirmRide)        //by User
		r.PATCH("/ridestatus/:userid", controllers.UpdateRideStatus) //by driver
		r.POST("/getridefare", controllers.GetRideFare)
	}

}
