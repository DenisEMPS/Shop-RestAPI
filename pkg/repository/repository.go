package repository

import (
	"school21_project1/types"

	"github.com/jmoiron/sqlx"
)

type Client interface {
	Create(client types.CreateClient) (int, error)
	Delete(id int) error
	Find(name string, surname string) ([]types.ClientDAO, error)
	GetAll(limit int, offset int) ([]types.ClientDAO, error)
	Update(id int, adress types.Adress) error
}

type Product interface {
	Create(product types.Product) (int, error)
	GetByID(id int) (types.ProductDAO, types.Image, error)
	GetAll(offset int, limit int) ([]types.ProductDAO, []types.Image, error)
	Delete(id int) error
	Update(id int, productU types.ProductUpdate) error
}

type Supplier interface {
	Create(supplier types.CreateSupplier) (int, error)
	Update(id int, adress types.Adress) error
	Delete(id int) error
	GetAll() ([]types.SupplierDTO, error)
	GetByID(id int) (types.SupplierDTO, error)
}

type Image interface {
	Create(image types.CreateImageProduct) (string, error)
	GetByID(id string) (types.Image, error)
	Update(uuid string, image types.Image) error
	Delete(uuid string) error
	GetByProductID(id int) (types.Image, error)
}

type Repositry struct {
	Client
	Product
	Supplier
	Image
}

func NewRepository(db *sqlx.DB) *Repositry {
	return &Repositry{
		Client:   NewClientPostgres(db),
		Product:  NewProductPostgres(db),
		Supplier: NewSupplierPostgres(db),
		Image:    NewImagePostgres(db),
	}
}
