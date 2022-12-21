package routers

import (
	"github.com/gin-gonic/gin"
	apiControllerV1 "man-CRM/controllers/api/v1"
	apiControllerV2 "man-CRM/controllers/api/v2"
	"man-CRM/middlewares"
)

// SetupRouter function will perform all route operations
func SetupRouter() *gin.Engine {

	route := gin.Default()

	//Giving access to storage folder
	route.Static("/storage", "storage")

	//Giving access to template folder
	route.Static("/templates", "templates")
	route.LoadHTMLGlob("templates/*")

	route.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	// load landing pages
	WebRoutes(route)

	//API route for version 1
	v1 := route.Group("/api/v1")
	v1.GET("test-response", apiControllerV1.TestResponse)
	//If you want to pass your route through specific middlewares
	v1.Use(middlewares.UserMiddlewares())
	{
		v1.POST("user-list", apiControllerV1.UserList)
	}

	//API route for version 2
	v2 := route.Group("/api/v2")

	v2.POST("user-list", apiControllerV2.UserList)

	return route

}
