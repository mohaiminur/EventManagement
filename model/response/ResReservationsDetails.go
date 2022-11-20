package response

type ResReservationsDetails struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	WorkshopId int    `json:"workshop_id,omitempty"`
}
