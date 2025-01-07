package service

import (
	"school21_project1/pkg/repository"
	"school21_project1/types"
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

type Service struct {
	Client
	Product
	Supplier
	Image
}

func NewService(repos *repository.Repositry) *Service {
	return &Service{
		Client:  NewClientService(repos.Client),
		Product: NewProductService(repos.Product),
	}
}
