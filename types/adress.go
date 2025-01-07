package types

type Adress struct {
	ID      string  `json:"-"`
	Country *string `json:"country"`
	City    *string `json:"city"`
	Street  *string `json:"street"`
}
