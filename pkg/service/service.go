package service

import (
	"school21_project1/pkg/repository"
	"school21_project1/types"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Client interface {
	Create(client types.CreateClient) (int, error)
	Delete(id int) error
	Find(name string, surname string) ([]types.ClientDAO, error)
	GetAll(limit int, offset int) ([]types.ClientDAO, error)
	Update(id int, adress types.AdressDTO) error
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
	Update(id int, adress types.AdressDTO) error
	Delete(id int) error
	GetAll() ([]types.SupplierDAO, error)
	GetByID(id int) (types.SupplierDAO, error)
}

type Image interface {
	Create(image types.CreateImageProduct) (string, error)
	GetByID(id string) (types.Image, error)
	Update(uuid string, image types.Image) error
	Delete(uuid string) error
	GetByProductID(id int) (types.Image, error)
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
		Image:    NewImageService(repos.Image),
	}
}
