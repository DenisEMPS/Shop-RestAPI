package types

type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Price          float64 `json:"price"`
	AvailableStock int     `json:"available_stock"`
	LastUpdateDate string  `json:"last_update_date"`
	SupplierID     string  `json:"supplier_id"`
	ImageID        string  `json:"image_id"`
}

type ProductDAO struct {
	ID                    int     `json:"-"`
	Name                  string  `json:"name" binding:"required"`
	Category              string  `json:"category" binding:"required"`
	Price                 float64 `json:"price" binding:"required"`
	AvailableStock        int     `json:"available_stock" binding:"required"`
	LastUpdateDate        string  `json:"last_update_date" binding:"required"`
	SupplierName          string  `json:"supplier_name" binding:"required"`
	SupplierAdressCountry string  `json:"adress_country" binding:"required"`
	SupplierAdressCity    string  `json:"adress_city" binding:"required"`
	SupplierAdressStreet  string  `json:"adress_street" binding:"required"`
	SupplierPhoneNumber   string  `json:"phone_number" binding:"required"`
	Image                 string  `json:"image" binding:"required"`
}
