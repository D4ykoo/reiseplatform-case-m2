package dto

type CreateHotelRequest struct {
	HotelName   string           `json:"hotelname"`
	Street      string           `json:"street"`
	State       string           `json:"state"`
	Land        string           `json:"land"`
	Description string           `json:"description"`
	Pictures    []PictureRequest `json:"pictures"`
}

type PictureRequest struct {
	Description string `json:"description"`
	Payload     string `json:"payload"`
}

type UpdateHotelRequest struct {
}

type CreateTravelRequest struct {
}

type DeleteTravelRequest struct {
}
