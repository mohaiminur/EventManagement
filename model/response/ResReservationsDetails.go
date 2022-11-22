package response

type ResReservationsDetails struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	WorkshopId int    `json:"workshop_id,omitempty"`
}

func (ResReservationsDetails) TableName() string {
	return "reservations"
}
