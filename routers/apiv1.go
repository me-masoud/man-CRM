package routers

import (
	"github.com/gin-gonic/gin"
	v1 "man-CRM/controllers/api/v1"
)

func ApiV1(version1 *gin.Engine) {
	Route := version1.Group("/api/v1")

	// Service
	Route.GET("services", v1.ServiceIndex)
	Route.GET("services/:id", v1.ServiceShow)
	Route.POST("services", v1.ServiceCreate)

	// Service Station
	Route.GET("service-stations", v1.ServiceStationIndex)
	Route.GET("service-stations/:id", v1.ServiceStationShow)
	Route.POST("service-stations", v1.ServiceStationCreate)

	// Stuff
	Route.GET("stuffs", v1.StuffIndex)
	Route.GET("stuffs/:id", v1.StuffShow)
	Route.POST("stuffs", v1.StuffCreate)

	// Status
}
