package response

type ResWorkshopsByEvents struct {
	Id        int                  `json:"id"`
	Title     string               `json:"title"`
	StartAt   string               `json:"start_at"`
	EndAt     string               `json:"end_at"`
	Workshops []ResWorkshopDetails `json:"workshops"`
	//EventDetails ResEventDetails      `json:"event_details"`
}

//	Workshops      []ResWorkshopDetails `json:"workshops,omitempty" bson:"workshops,omitempty"`
//	Workshops      []ResWorkshopDetails `json:"workshops,omitempty"`
//type ResEventDetails struct {
//	ID        string    `json:"id"`
//	Title     string    `json:"title"`
//	StartAt   string    `json:"start_at"`
//	EndAt     time.Time `json:"end_at"`
//	Workshops []struct {
//		ID          string `json:"id"`
//		Title       string `json:"title"`
//		Description string `json:"description"`
//		StartAt     string `json:"start_at"`
//		EndAt       string `json:"end_at"`
//	} `json:"workshops"`
//}

//type ResEventDetails struct {
//	Id             int    `json:"id"`
//	Title          string `json:"title"`
//	StartAt        string `json:"start_at"`
//	EndAt          string `json:"end_at"`
//	TotalWorkshops int    `json:"total_workshops,omitempty"`
//}

//	Id          int    `json:"id,omitempty"`
