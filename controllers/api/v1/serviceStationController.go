package v1

import (
	"github.com/gin-gonic/gin"
	"man-CRM/models"
	"net/http"
)

func ServiceStationIndex(c *gin.Context) {

	services := models.GetDB().Offset(3).Find(&models.ServiceStation{})
	c.JSON(http.StatusOK, services)
}

func ServiceStationCreate(c *gin.Context) {
	//status := models.GetDB().First(&models.Status{})

	service := models.GetDB().Create(&models.ServiceStation{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
	})
	c.JSON(http.StatusOK, service.Value)

}

func ServiceStationShow(c *gin.Context) {
	service := models.GetDB().First(&models.ServiceStation{}, c.Param("id"))
	c.JSON(http.StatusOK, service)
}
