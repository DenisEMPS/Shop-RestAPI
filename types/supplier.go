package types

type CreateSupplier struct {
	Name        string `json:"name" binding:"required"`
	Country     string `json:"country" binding:"required"`
	City        string `json:"city" binding:"required"`
	Street      string `json:"street" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type SupplierDTO struct {
	Name        string `json:"name" binding:"required" db:"supplier_name"`
	Country     string `json:"country" binding:"required" db:"country"`
	City        string `json:"city" binding:"required" db:"city"`
	Street      string `json:"street" binding:"required" db:"street"`
	PhoneNumber string `json:"phone_number" binding:"required" db:"supplier_phone_number"`
}
