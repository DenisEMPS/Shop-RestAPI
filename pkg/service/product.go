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

func (p ProductService) Create(product types.ProductDAO) (int, error) {
	return p.repo.Create(product)
}
