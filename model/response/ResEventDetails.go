package response

type ResEventDetails struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	StartAt        string `json:"start_at"`
	EndAt          string `json:"end_at"`
	TotalWorkshops int    `json:"total_workshops,omitempty"`
}
type ResEventDetailsWithPag struct {
	ResEventDetails []ResEventDetails `json:"id"`
	Pagination      Pagination        `json:"pagination"`
}
type Pagination struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
}

//	Workshops      []ResWorkshopDetails `json:"workshops,omitempty" bson:"workshops,omitempty"`
//	Workshops      []ResWorkshopDetails `json:"workshops,omitempty"`
