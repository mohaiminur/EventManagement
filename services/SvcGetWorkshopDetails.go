package services

import (
	"EventManagement/config"
	"EventManagement/model/response"
)

var WorkshopDetails []response.ResWorkshopDetails

func SvcGetWorkshopDetails(id string) ([]response.ResWorkshopDetails, error) {

	config.DB.Raw("select w.id, w.title, w.description, w.start_at, w.end_at, count(r.id) as total_reservations "+
		"from workshops w left join reservations r on w.id = r.workshop_id where w.id= ?", id).Scan(&WorkshopDetails)

	return WorkshopDetails, nil
}

var WorkshopDetail response.WorkshopDetails

func GetWorkshopDetails(id string) (response.WorkshopDetails, error) {

	config.DB.Raw("select w.id, w.event_id, w.title, w.description, w.start_at, w.end_at "+
		"from workshops w where w.id= ?", id).Scan(&WorkshopDetail)

	return WorkshopDetail, nil
}
