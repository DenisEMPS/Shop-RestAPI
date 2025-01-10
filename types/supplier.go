package types

type CreateSupplier struct {
	Name        string `json:"name" binding:"required"`
	Country     string `json:"country" binding:"required"`
	City        string `json:"city" binding:"required"`
	Street      string `json:"street" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type SupplierDAO struct {
	Name        string `json:"name" db:"supplier_name"`
	Country     string `json:"country" db:"country"`
	City        string `json:"city" db:"city"`
	Street      string `json:"street" db:"street"`
	PhoneNumber string `json:"phone_number" db:"supplier_phone_number"`
}
