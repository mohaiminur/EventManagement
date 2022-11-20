package handler

import (
	"EventManagement/model/response"
	"EventManagement/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWorkshopDetails(c *gin.Context) {
	id := c.Param("workshopId")
	WorkshopDetails, err := services.SvcGetWorkshopDetails(id)
	if err != nil {
		panic(err)
	}
	var WorkshopDetail response.ResWorkshopDetails
	if WorkshopDetails.Id == 0 {
		c.JSON(404, "Workshop not found by id: "+id)
	} else {
		WorkshopDetail = WorkshopDetails

		c.JSON(http.StatusOK, WorkshopDetail)
	}
}
