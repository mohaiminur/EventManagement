package main

import (
	"EventManagement/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	//router.GET("/event-list/:limit/:offset", handler.GetEvents)

	router.GET("/event-list/:limit/:offset", handler.GetEvents)
	router.GET("get-event-details/:eventId", handler.GetEventDetails)
	router.GET("get-workshop-details/:workshopId", handler.GetWorkshopDetails)
	router.GET("get-workshop-details-by-event/:eventId", handler.GetWorkshopByEvent)
	router.GET("workshop-reservation", handler.WorkshopReservation)

	//router.GET("/ugc-terms-and-conditions/:deviceType/:dbVersion/:theme", middleware.Details, handler.Details)

}
