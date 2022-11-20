package response

type ResWorkshopsByEvents struct {
	Id        int                  `json:"id"`
	Title     string               `json:"title"`
	StartAt   string               `json:"start_at"`
	EndAt     string               `json:"end_at"`
	Workshops []ResWorkshopDetails `json:"workshops"`
}
