package v1

import (
	"github.com/gin-gonic/gin"
	"man-CRM/models"
	"net/http"
)

func ServiceIndex(c *gin.Context) {

	services := models.GetDB().Offset(3).Find(&models.Service{})
	c.JSON(http.StatusOK, services)
}

func ServiceCreate(c *gin.Context) {
	//status := models.GetDB().First(&models.Status{})

	service := models.GetDB().Create(&models.Service{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
	})
	c.JSON(http.StatusOK, service.Value)

}
