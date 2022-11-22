package services

import (
	"EventManagement/config"
	"EventManagement/model/request"
	"EventManagement/model/response"
	"strconv"
)

func SvcGetReservationDetails(name string, email string) (response.ResReservationsDetails, error) {
	var ReservationDetails response.ResReservationsDetails
	config.DB.Raw("select r.id, r.name, r.email, r.workshop_id " +
		"from reservations r  where name= '" + name + "'" + " and " + " email= '" + email + "'").Scan(&ReservationDetails)
	return ReservationDetails, nil
}

func CreateReservation(body request.WorkshopReservationBody) (response.ResReservationsDetails, error) {
	var ReservationDetails1 response.ResReservationsDetails
	ReservationDetails1.Name = body.Name
	ReservationDetails1.Email = body.Email
	ReservationDetails1.WorkshopId, _ = strconv.Atoi(body.WorkshopId)
	config.DB.Create(&ReservationDetails1)
	//	config.DB.Create(&body).Scan(ReservationDetails1)
	return ReservationDetails1, nil
}

func GetNameEmail(name string, email string) (response.ResReservationsDetails, error) {
	var ReservationDetails response.ResReservationsDetails
	config.DB.Raw("select r.id, r.name, r.email, r.workshop_id " +
		"from reservations r  where name= '" + name + "'" + " or " + " email= '" + email + "'").Scan(&ReservationDetails)
	return ReservationDetails, nil
}
