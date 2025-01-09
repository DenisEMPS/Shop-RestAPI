package service

import (
	"school21_project1/pkg/repository"
	"school21_project1/types"
)

type ImageService struct {
	repo repository.Image
}

func NewImageService(repo repository.Image) *ImageService {
	return &ImageService{repo: repo}
}

func (i ImageService) Create(image types.CreateImageProduct) (string, error) {
	return i.repo.Create(image)
}

func (i ImageService) GetByID(id string) (types.Image, error) {
	return i.repo.GetByID(id)
}

func (i ImageService) Update(uuid string, image types.Image) error {
	return i.repo.Update(uuid, image)
}

func (i ImageService) Delete(uuid string) error {
	return i.repo.Delete(uuid)
}

func (i ImageService) GetByProductID(id int) (types.Image, error) {
	return i.repo.GetByProductID(id)
}
