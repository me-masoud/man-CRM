package routers

import (
	"github.com/gin-gonic/gin"
	"man-CRM/controllers/api/v1/HomeController"
)

func WebRoutes(route *gin.Engine) {
	route.GET("home", HomeController.Index)
}
