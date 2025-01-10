package types

type AdressDTO struct {
	Country *string `json:"country"`
	City    *string `json:"city"`
	Street  *string `json:"street"`
}
