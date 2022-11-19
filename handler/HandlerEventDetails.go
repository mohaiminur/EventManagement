package handler

import (
	"EventManagement/model/response"
	"EventManagement/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
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
func GetEvents(c *gin.Context) {
	limit := c.Param("limit")
	offset := c.Param("limit")
	//limit :=  strconv.Atoi(limitt)
	//offset :=  strconv.Atoi(offsett)

	fmt.Println(limit + offset)

	EventDetails, err := services.GetEvents(limit, offset)
	if err != nil {
		panic(err)
	}
	fmt.Println(EventDetails)
	//var eventDetail response.ResEventDetails
	//if EventDetails.Id == 0 {
	//	c.JSON(404, "Event not found by id: "+id)
	//} else {
	//	eventDetail = EventDetails
	//
	c.JSON(http.StatusOK, EventDetails)
}
