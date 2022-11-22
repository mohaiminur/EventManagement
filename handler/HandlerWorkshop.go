package handler

import (
	"EventManagement/model/request"
	"EventManagement/model/response"
	"EventManagement/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var t = time.Now()
var now = t.Format("2006-01-02 15:04:05")

func GetWorkshopByEvent(c *gin.Context) {

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
	fmt.Println(reservation.WorkshopId)
	if reservation.Id == 0 {
		c.JSON(400, "name or email is not correct")
	} else {
		c.JSON(http.StatusOK, resWorkshopReservation)
	}

}

func WorkshopReservationV2(c *gin.Context) {
	//first get workshop details from wid body name email
	//find reservation from this id or not
	//
	var resReservation response.ResReservationsDetails
	var resWorkshop response.WorkshopDetails

	var workResBody request.WorkshopReservationBody
	if err := c.ShouldBindJSON(&workResBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation, err1 := services.SvcGetReservationDetails(workResBody.Name, workResBody.Email)

	////if not exist name and email
	if err1 != nil {
		c.JSON(400, "failed to reservation from email and name  in front ")
		//|| reservation.WorkshopId == 0 || &reservation.Name == nil || &reservation.Email == nil
	} else if reservation.Id == 0 {
		reservation, err1 := services.GetNameEmail(workResBody.Name, workResBody.Email)
		if err1 != nil {
			fmt.Println("failed to find")
		}
		if reservation.Name == workResBody.Name {
			c.IndentedJSON(400, " User already exist!! try to create reservation in different name or enter correct email")
			c.Abort()
		} else if reservation.Email == workResBody.Email {
			c.IndentedJSON(400, " Email already exist!! try to create reservation ith different email or enter correct name")
			c.Abort()
		} else {

			//check if workshop is present or not from post body workshop id
			workshopIsPresentOrNot, err := services.GetWorkshopDetail(workResBody.WorkshopId)
			if err != nil {
				fmt.Println("failed to find workshop by id ")
			}
			if workshopIsPresentOrNot.Id == 0 {
				workshopIdList, _ := services.SvcGetWorkshopIdList(now)

				var wd string
				wd = "{ "
				for i := 0; i <= len(workshopIdList)-1; i++ {
					wd += strconv.Itoa(workshopIdList[i].Id) + ", "
				}
				wd += "}"

				fmt.Println(wd)
				c.IndentedJSON(400, "To create reservation please select workshop from the available workshop list : "+wd)
				c.Abort()
				//c.Abort()
			} else {
				//creating a reservation
				createRes, err1 := services.CreateReservation(workResBody)
				if err1 != nil {
					c.JSON(400, "failed to reservation from email and name after creation  in create sec ")
				} else if createRes.Name == workResBody.Name {

					//reservation, err1 := services.SvcGetReservationDetails(workResBody.Name, workResBody.Email)
					//if err1 != nil {
					//	c.JSON(400, "failed to reservation from email and name after creation  in get sec ")
					//}
					resReservation.Id = createRes.Id
					resReservation.Name = createRes.Name
					resReservation.Email = createRes.Email

					workshop, err := services.GetWorkshopDetail(workResBody.WorkshopId)
					if err != nil {
						c.JSON(400, "failed to find workshop in create sec ")
					}
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
					if createRes.Id == 0 {
						c.JSON(400, "name or email is not correct in create get sec")
					} else {
						c.JSON(http.StatusOK, resWorkshopReservation)
					}
				}

			}
		}

		//}
		//// if already exist name and email----------------------------------------------
	} else {

		reservation, err1 := services.SvcGetReservationDetails(workResBody.Name, workResBody.Email)
		if err1 != nil {
			c.JSON(400, "failed to reservation from email and name after creation  in get sec ")
			c.Abort()
		}
		workshop, err := services.GetWorkshopDetail(strconv.Itoa(reservation.WorkshopId))
		if err != nil {
			c.JSON(400, "failed to get workshop id  in get sec ")
			c.Abort()
		}
		resWorkshop.Id = workshop.Id
		resWorkshop.Title = workshop.Title
		resWorkshop.Description = workshop.Description
		resWorkshop.StartAt = workshop.StartAt
		resWorkshop.EndAt = workshop.EndAt

		eventid := workshop.EventId

		//gotevent
		event, err := services.GetEventDetails(eventid)

		resReservation.Id = reservation.Id
		resReservation.Name = reservation.Name
		resReservation.Email = reservation.Email

		var resWorkshopReservation response.ResWorkshopReservation
		resWorkshopReservation.Reservation = resReservation
		resWorkshopReservation.Event = event
		resWorkshopReservation.Workshop = resWorkshop

		if err != nil {
			panic(err)
		}
		if reservation.Id == 0 {
			c.JSON(400, "name or email is not correct  in get sec last")
			c.Abort()
		} else {
			c.JSON(http.StatusOK, resWorkshopReservation)
		}
	}

}
