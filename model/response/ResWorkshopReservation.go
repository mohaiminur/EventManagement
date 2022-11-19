package response

type ResWorkshopReservation struct {
	Reservation ResReservationsDetails `json:"reservation"`
	Event       ResEventDetails        `json:"event"`
	Workshop    WorkshopDetails        `json:"workshop"`
}
