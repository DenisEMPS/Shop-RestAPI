package service

import (
	"school21_project1/pkg/repository"
	"school21_project1/types"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (p ProductService) Create(product types.Product) (int, error) {
	return p.repo.Create(product)
}

func (p ProductService) GetByID(id int) (types.ProductDAO, error) {
	return p.repo.GetByID(id)
}

func (p ProductService) GetAll(offset int, limit int) ([]types.ProductDAO, error) {
	return p.repo.GetAll(offset, limit)
}

func (p ProductService) Delete(id int) error {
	return p.repo.Delete(id)
}

func (p ProductService) Update(id int, productU types.ProductUpdate) error {
	return p.repo.Update(id, productU)
}
