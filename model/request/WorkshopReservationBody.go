package request

type WorkshopReservationBody struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	WorkshopId string `json:"workshop_id"`
}
