package handler

import (
	"EventManagement/model/request"
	"EventManagement/model/response"
	"EventManagement/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetWorkshopByEvent(c *gin.Context) {

	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")

	id := c.Param("eventId")
	var resWorkshopByEv response.ResWorkshopsByEvents
	EventDetails, err := services.SvcGetEventDetails(id)

	Workshop, err := services.SvcGetWorkshopByEvent(id, now)
	if err != nil {
		panic(err)
	}

	resWorkshopByEv.Id = EventDetails.Id
	resWorkshopByEv.Title = EventDetails.Title
	resWorkshopByEv.StartAt = EventDetails.StartAt
	resWorkshopByEv.EndAt = EventDetails.EndAt
	resWorkshopByEv.Workshops = Workshop

	if EventDetails.Id == 0 {
		c.JSON(404, "Event not found by id: "+id)
	} else {

		c.JSON(http.StatusOK, resWorkshopByEv)
	}

}
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

func WorkshopReservation(c *gin.Context) {

	var workResBody request.WorkshopReservationBody
	if err := c.ShouldBindJSON(&workResBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var resReservation response.ResReservationsDetails
	var resWorkshop response.WorkshopDetails
	//var resEvent response.ResEventDetails

	//got reservation
	reservation, err := services.SvcGetReservationDetails(workResBody.Name, workResBody.Email)

	resReservation.Id = reservation.Id
	resReservation.Name = reservation.Name
	resReservation.Email = reservation.Email

	workshopid := strconv.Itoa(reservation.WorkshopId)
	//got workshop

	workshop, err := services.GetWorkshopDetail(workshopid)
	resWorkshop.Id = workshop.Id
	resWorkshop.Title = workshop.Title
	resWorkshop.Description = workshop.Description
	resWorkshop.StartAt = workshop.StartAt
	resWorkshop.EndAt = workshop.EndAt

	eventid := workshop.EventId

	//gotevent
	event, err := services.GetEventDetails(eventid)

	var resWorkshopReservation response.ResWorkshopReservation
	resWorkshopReservation.Reservation = resReservation
	resWorkshopReservation.Event = event
	resWorkshopReservation.Workshop = resWorkshop

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, resWorkshopReservation)

}
