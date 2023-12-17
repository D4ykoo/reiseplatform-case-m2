package dto

type TravelResponse struct {
}

type HotelRespnse struct {
	Id          string            `json:"id"`
	HotelName   string            `json:"hotelname"`
	Street      string            `json:"street"`
	State       string            `json:"state"`
	Land        string            `json:"land"`
	VendorID    string            `json:"vendorid"`
	VendorName  string            `json:"vendorname"`
	Description string            `json:"description"`
	Pictures    []PictureResponse `json:"pictures"`
}

type PictureResponse struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Payload     string `json:"payload"`
}
