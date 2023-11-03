package router

import (
	"taxi_app/controllers"
	"taxi_app/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {
	r := router.Group("/admin")
	r.POST("/login", controllers.AdminLogin)
	r.GET("/pendingdrivers", middleware.AdminAuth, controllers.GetPendingDrivers)
	r.PATCH("/acceptdriver/:id", middleware.AdminAuth, controllers.AcceptDrivers)
	r.PATCH("/rejectdriver/:id", middleware.AdminAuth, controllers.RejectDrivers)
	r.GET("/accepteddrivers", middleware.AdminAuth, controllers.GetAcceptedDrivers)
	r.GET("/rejecteddrivers", middleware.AdminAuth, controllers.GetRejectedDrivers)
	r.POST("/coupon", middleware.AdminAuth, controllers.CreateCoupons)
	r.GET("/ongingcoupons", middleware.AdminAuth, controllers.GetOngoingCoupons)
}
