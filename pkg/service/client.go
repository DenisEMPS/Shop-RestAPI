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

func (c ClientService) Create(client types.Client) (int, error) {
	return c.repo.Create(client)
}

func (c ClientService) Delete(id int) error {
	return c.repo.Delete(id)
}

func (c ClientService) Find(name string, surname string) ([]types.Client, error) {
	return c.repo.Find(name, surname)
}

func (c ClientService) GetAll(limit string, offset string) ([]types.ClientDTO, error) {
	return c.repo.GetAll(limit, offset)
}

func (c ClientService) Update(id string, adress types.Adress) error {
	return c.repo.Update(id, adress)
}
