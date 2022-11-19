package response

type ResWorkshopDetails struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	StartAt           string `json:"start_at"`
	EndAt             string `json:"end_at"`
	TotalReservations int    `json:"total_reservations,omitempty"`
}
type WorkshopDetails struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EventId     string `json:"event_id,omitempty"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
}
