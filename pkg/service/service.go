package service

import (
	"school21_project1/pkg/repository"
	"school21_project1/types"
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
	Create(supplier types.CreateSupplier) (int, error)
	Update(id int, adress types.Adress) error
	Delete(id int) error
	GetAll() ([]types.SupplierDTO, error)
	GetByID(id int) (types.SupplierDTO, error)
}

type Image interface {
}

type Service struct {
	Client
	Product
	Supplier
	Image
}

func NewService(repos *repository.Repositry) *Service {
	return &Service{
		Client:   NewClientService(repos.Client),
		Product:  NewProductService(repos.Product),
		Supplier: NewSupplierService(repos.Supplier),
	}
}
