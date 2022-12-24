package v1

import (
	"github.com/gin-gonic/gin"
	"man-CRM/models"
	"net/http"
)

func StuffIndex(c *gin.Context) {

	services := models.GetDB().Offset(3).Find(&models.Stuff{})
	c.JSON(http.StatusOK, services)
}

func StuffCreate(c *gin.Context) {
	//status := models.GetDB().First(&models.Status{})

	service := models.GetDB().Create(&models.Stuff{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
	})
	c.JSON(http.StatusOK, service.Value)

}

func StuffShow(c *gin.Context) {
	service := models.GetDB().First(&models.Stuff{}, c.Param("id"))
	c.JSON(http.StatusOK, service)
}
