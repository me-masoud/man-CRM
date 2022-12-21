package HomeController

import (
	"github.com/gin-gonic/gin"
	u "man-CRM/apiHelpers"
)

func Index(c *gin.Context) {
	u.Respond(c.Writer, u.Message(200, "This is Home"))
}
