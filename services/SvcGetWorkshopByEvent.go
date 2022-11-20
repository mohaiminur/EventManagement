package services

import (
	"EventManagement/config"
	"EventManagement/model/response"
)

var WorkshopDetailsByEvent []response.ResWorkshopDetails

func SvcGetWorkshopByEvent(id string, now string) ([]response.ResWorkshopDetails, error) {

	config.DB.Raw("select w.id, w.title, w.description, w.start_at, w.end_at"+
		" from workshops w  where w.start_at > '"+now+"' and w.event_id= ?", id).Scan(&WorkshopDetailsByEvent)

	return WorkshopDetailsByEvent, nil
}
