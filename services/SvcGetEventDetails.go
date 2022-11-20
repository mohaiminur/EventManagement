package services

import (
	"EventManagement/config"
	"EventManagement/model/response"
)

var AllEventDetails []response.ResEventDetails

func GetEvents(limit string, offset string, now string) ([]response.ResEventDetails, error) {

	config.DB.Raw("select * from events e where e.start_at > '" + now + "' LIMIT " + limit + " offset " + offset + "").Scan(&AllEventDetails)
	//	fmt.Println("select * from events e  LIMIT " + limit + " offset " + offset + "")
	return AllEventDetails, nil
}

var EventDetails response.ResEventDetails

func GetTotal(now string) (response.ResEventDetails, error) {
	config.DB.Raw("select count(id) as id from events e where e.start_at > '" + now + "' ").Scan(&EventDetails.Id)

	return EventDetails, nil
}

func SvcGetEventDetails(id string) (response.ResEventDetails, error) {

	config.DB.Raw("select e.id, e.title, e.start_at, e.end_at, count(w.id) as total_workshops from events e left join workshops w on e.id = w.event_id where e.id= ?", id).Scan(&EventDetails)

	return EventDetails, nil
}
func GetEventDetails(id string) (response.ResEventDetails, error) {

	config.DB.Raw("select e.id, e.title, e.start_at, e.end_at from events e  where e.id= ?", id).Scan(&EventDetails)

	return EventDetails, nil
}
