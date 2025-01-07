package repository

import (
	"school21_project1/types"

	"github.com/jmoiron/sqlx"
)

type Client interface {
	Create(client types.CreateClient) (int, error)
	Delete(id int) error
	Find(name string, surname string) ([]types.ClientDTO, error)
	GetAll(limit int, offset int) ([]types.ClientDTO, error)
	Update(id int, adress types.Adress) error
}

type Product interface {
	Create(product types.Product) (int, error)
	GetByID(id int) (types.ProductDAO, error)
	GetAll(offset int, limit int) ([]types.ProductDAO, error)
	Delete(id int) error
	Update(id int, productU types.ProductUpdate) error
}

type Supplier interface {
}

type Image interface {
}

type Repositry struct {
	Client
	Product
	Supplier
	Image
}

func NewRepository(db *sqlx.DB) *Repositry {
	return &Repositry{
		Client:  NewClientPostgres(db),
		Product: NewProductPostgres(db),
	}
}
