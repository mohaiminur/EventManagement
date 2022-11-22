package services

import (
	"EventManagement/config"
	"EventManagement/model/response"
	"fmt"
)

func SvcGetWorkshopDetails(id string) (response.ResWorkshopDetails, error) {
	var WorkshopDetails response.ResWorkshopDetails
	config.DB.Raw("select w.id, w.title, w.description, w.start_at, w.end_at, count(r.id) as total_reservations "+
		"from workshops w left join reservations r on w.id = r.workshop_id where w.id= ?", id).Scan(&WorkshopDetails)
	return WorkshopDetails, nil
}

func GetWorkshopDetail(id string) (response.WorkshopDetails, error) {
	var WorkshopDetail response.WorkshopDetails
	config.DB.Raw("select w.id, w.event_id, w.title, w.description, w.start_at, w.end_at "+
		"from workshops w where w.id= ?", id).Scan(&WorkshopDetail)
	return WorkshopDetail, nil
}

func SvcGetWorkshopByEvent(id string, now string) ([]response.ResWorkshopDetails, error) {
	var WorkshopDetailsByEvent []response.ResWorkshopDetails
	config.DB.Raw("select w.id, w.title, w.description, w.start_at, w.end_at"+
		" from workshops w  where w.start_at > '"+now+"' and w.event_id= ?", id).Scan(&WorkshopDetailsByEvent)
	return WorkshopDetailsByEvent, nil
}

func SvcGetWorkshopIdList(now string) ([]response.ResWorkshopDetails, error) {
	var WorkshopDetailsByEvent []response.ResWorkshopDetails
	config.DB.Raw("select w.id  from workshops w  where w.start_at > '" + now + "'").Scan(&WorkshopDetailsByEvent)
	fmt.Println("select w.id  from workshops w  where w.start_at > '" + now + "'")
	return WorkshopDetailsByEvent, nil
}
