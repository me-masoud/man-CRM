package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	u "man-CRM/apiHelpers"
	v1s "man-CRM/services/api/v1"
)

func ServiceIndex(c *gin.Context) {
	var userService v1s.UserService

	//decode the request body into struct and failed if any error occur
	err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	if err != nil {
		u.Respond(c.Writer, u.Message(1, "Invalid request"))
		return
	}

	//call service
	resp := userService.UserList()

	//return response using api helper
	u.Respond(c.Writer, resp)
}

func ServiceCreate(c *gin.Context) {

}
