package routers

import (
	"github.com/gin-gonic/gin"
	v1 "man-CRM/controllers/api/v1"
)

func ApiV1(version1 *gin.Engine) {
	Route := version1.Group("/api/v1")
	Route.GET("services", v1.ServiceIndex)
}
