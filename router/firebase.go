package router

import (
	"taxi_app/controllers"

	"github.com/gin-gonic/gin"
)

func FirebaseRouter(router *gin.Engine) {
	r := router.Group("/firebase")
	r.POST("/token", controllers.FirebaseCredentials)
	r.PATCH("status",controllers.UpdateStatus)
}
