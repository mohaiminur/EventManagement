package handler

import (
	"EventManagement/model/response"
	"EventManagement/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEventDetails(c *gin.Context) {
	id := c.Param("eventId")
	EventDetails, err := services.SvcGetEventDetails(id)
	if err != nil {
		panic(err)
	}
	var eventDetail response.ResEventDetails
	if EventDetails.Id == 0 {
		c.JSON(404, "Event not found by id: "+id)
	} else {
		eventDetail = EventDetails

		c.JSON(http.StatusOK, eventDetail)
	}

}
