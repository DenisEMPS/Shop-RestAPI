package types

type CreateImageProduct struct {
	Image     []byte `json:"image"`
	ProductID int    `json:"product_id"`
}

type Image struct {
	Image []byte `json:"image" db:"image"`
}
