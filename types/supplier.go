package types

type Supplier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AddressID   string `json:"adress_id"`
	PhoneNumber string `json:"phone_number"`
}
