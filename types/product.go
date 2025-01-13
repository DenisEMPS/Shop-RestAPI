package types

type Product struct {
	Name           string  `json:"name" binding:"required"`
	Category       string  `json:"category" binding:"required"`
	Price          float64 `json:"price" binding:"required"`
	AvailableStock int     `json:"available_stock" binding:"required"`
	LastUpdateDate string  `json:"last_update_date" binding:"required"`
	SupplierID     int     `json:"supplier_id" binding:"required"`
	Image          []byte  `json:"image,omitempty"`
}

type ProductDAO struct {
	Name                  string  `json:"name" db:"name"`
	Category              string  `json:"category" db:"category"`
	Price                 float64 `json:"price" db:"price"`
	AvailableStock        int     `json:"available_stock" db:"available_stock"`
	LastUpdateDate        string  `json:"last_update_date" db:"last_update_date"`
	SupplierName          string  `json:"supplier_name" db:"supplier_name"`
	SupplierAdressCountry string  `json:"adress_country" db:"country"`
	SupplierAdressCity    string  `json:"adress_city" db:"city"`
	SupplierAdressStreet  string  `json:"adress_street" db:"street"`
	SupplierPhoneNumber   string  `json:"phone_number" db:"supplier_phone_number"`
}

type ProductUpdate struct {
	AvailableStockU int `json:"available_stock" binding:"required"`
}
