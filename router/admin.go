package router

import (
	"taxi_app/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	r.GET("/pendingdrivers", controllers.GetPendingDrivers)
	r.PATCH("/acceptdriver/:id", controllers.AcceptDrivers)
	r.PATCH("/rejectdriver/:id", controllers.RejectDrivers)
	r.GET("/accepteddrivers", controllers.GetAcceptedDrivers)
	r.GET("/rejecteddrivers", controllers.GetRejectedDrivers)
}
