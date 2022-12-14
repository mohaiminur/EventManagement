package handler

import (
	"EventManagement/model/response"
	"EventManagement/services"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"time"
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
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	//fmt.Println(t.Format(time.RFC3339))

	limitt := c.Param("limit")
	offsett := c.Param("offset")
	limit, _ := strconv.Atoi(limitt)
	offset, _ := strconv.Atoi(offsett)
	pages := limit

	PerPage := (offset - 1) * (limit)
	TotalEvent, _ := services.GetTotal(now)
	EventDetails, err := services.GetEvents(limitt, strconv.Itoa(PerPage), now)
	if err != nil {
		panic(err)
	}
	Total := TotalEvent.Id
	perPage := limit
	TotalPages := math.Ceil(float64(Total) / float64(pages))
	currentPages := offset

	var pagination response.Pagination
	pagination.Total = Total
	pagination.PerPage = perPage
	pagination.CurrentPage = currentPages
	pagination.TotalPages = int(TotalPages)

	var resEventDetailsWithPag response.ResEventDetailsWithPag
	resEventDetailsWithPag.ResEventDetails = EventDetails
	resEventDetailsWithPag.Pagination = pagination

	c.JSON(http.StatusOK, resEventDetailsWithPag)
}
