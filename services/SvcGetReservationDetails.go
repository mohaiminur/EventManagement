package services

import (
	"EventManagement/config"
	"EventManagement/model/response"
)

var ReservationDetails response.ResReservationsDetails

func SvcGetReservationDetails(name string, email string) (response.ResReservationsDetails, error) {

	config.DB.Raw("select r.id, r.name, r.email, r.workshop_id " +
		"from reservations r  where name= '" + name + "' and email= '" + email + "'").Scan(&ReservationDetails)

	//fmt.Println("select r.id, r.name, r.email, r.workshop_id " + "from reservations r  where name= " + name + " and email= " + email + "")
	return ReservationDetails, nil
}
