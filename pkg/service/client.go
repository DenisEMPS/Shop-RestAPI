package service

import (
	"school21_project1/pkg/repository"
	"school21_project1/types"
)

type ClientService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{repo: repo}
}

func (c ClientService) Create(client types.CreateClient) (int, error) {
	return c.repo.Create(client)
}

func (c ClientService) Delete(id int) error {
	return c.repo.Delete(id)
}

func (c ClientService) Find(name string, surname string) ([]types.ClientDAO, error) {
	return c.repo.Find(name, surname)
}

func (c ClientService) GetAll(limit int, offset int) ([]types.ClientDAO, error) {
	return c.repo.GetAll(limit, offset)
}

func (c ClientService) Update(id int, adress types.Adress) error {
	return c.repo.Update(id, adress)
}
