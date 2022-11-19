package handler

import (
	"EventManagement/model/response"
	"EventManagement/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWorkshopByEvent(c *gin.Context) {
	id := c.Param("eventId")
	var resWorkshopByEv response.ResWorkshopsByEvents
	EventDetails, err := services.SvcGetEventDetails(id)

	Workshop, err := services.SvcGetWorkshopByEvent(id)
	if err != nil {
		panic(err)
	}
	//Workshops := []response.ResWorkshopDetails{}
	//if len(Workshop) > 0 {
	//	for i := 0; i < len(Workshop); i++ {
	//		Workshops[i] = Workshop[i]
	//		//	workshops=append(workshops,workshop[i])
	//	}
	//}
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
