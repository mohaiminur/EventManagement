package handler

import (
	"EventManagement/model/request"
	"EventManagement/model/response"
	"EventManagement/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
