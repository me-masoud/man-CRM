package v1

import (
	"github.com/gin-gonic/gin"
	u "man-CRM/apiHelpers"
)

func TestResponse(c *gin.Context) {
	u.Respond(c.Writer, u.Message(200, "api works fine"))
}
