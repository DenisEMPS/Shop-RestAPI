package types

type Product struct {
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Price          float64 `json:"price"`
	AvailableStock int     `json:"available_stock"`
	LastUpdateDate string  `json:"last_update_date"`
	SupplierID     string  `json:"supplier_id"`
	Image          []byte  `json:"image,omitempty"`
}

type ProductDAO struct {
	ID                    int     `json:"-"`
	Name                  string  `json:"name" binding:"required" db:"name"`
	Category              string  `json:"category" binding:"required" db:"category"`
	Price                 float64 `json:"price" binding:"required" db:"price"`
	AvailableStock        int     `json:"available_stock" binding:"required" db:"available_stock"`
	LastUpdateDate        string  `json:"last_update_date" binding:"required" db:"last_update_date"`
	SupplierName          string  `json:"supplier_name" binding:"required" db:"suppler_name"`
	SupplierAdressCountry string  `json:"adress_country" binding:"required" db:"country"`
	SupplierAdressCity    string  `json:"adress_city" binding:"required" db:"city"`
	SupplierAdressStreet  string  `json:"adress_street" binding:"required" db:"street"`
	SupplierPhoneNumber   string  `json:"phone_number" binding:"required" db:"phone_number"`
	Image                 []byte  `json:"image" binding:"required" db:"image_id"`
}

type ProductUpdate struct {
	AvailableStockU int `json:"available_stock" binding:"required"`
}
