package router

import (
	"taxi_app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	r := router.Group("/user")
	{
		r.POST("/signup", controllers.SignUpUser)
		r.POST("/signin", controllers.LoginUser)
		r.PATCH("/payment", controllers.Payment)
		r.GET("/paymentdetails/:user_id", controllers.CompletedTrips)
		r.GET("/activetrip", controllers.ActiveTrips)
	}

}
