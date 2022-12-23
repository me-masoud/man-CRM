package routers

import (
	"github.com/gin-gonic/gin"
	v1 "man-CRM/controllers/api/v1"
)

func ApiV1(version1 *gin.Engine) {
	Route := version1.Group("/api/v1")

	// Service
	Route.GET("services", v1.ServiceIndex)
	Route.POST("services", v1.ServiceCreate)
	Route.GET("services/:id", v1.ServiceShow)
}
