package service

import (
	"school21_project1/pkg/repository"
	"school21_project1/types"
)

type SupplierService struct {
	repo repository.Supplier
}

func NewSupplierService(repo repository.Supplier) *SupplierService {
	return &SupplierService{repo: repo}
}

func (s SupplierService) Create(supplier types.CreateSupplier) (int, error) {
	return s.repo.Create(supplier)
}

func (s SupplierService) Update(id int, adress types.AdressDTO) error {
	return s.repo.Update(id, adress)
}

func (s SupplierService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s SupplierService) GetAll() ([]types.SupplierDAO, error) {
	return s.repo.GetAll()
}

func (s SupplierService) GetByID(id int) (types.SupplierDAO, error) {
	return s.repo.GetByID(id)
}
