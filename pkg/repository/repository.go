package repository

import (
	"school21_project1/types"

	"github.com/jmoiron/sqlx"
)

type Client interface {
	Create(client types.Client) (int, error)
	Delete(id int) error
	Find(name string, surname string) ([]types.Client, error)
	GetAll(limit string, offset string) ([]types.ClientDTO, error)
	Update(id string, adress types.Adress) error
}

type Product interface {
	Create(product types.ProductDAO) (int, error)
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
