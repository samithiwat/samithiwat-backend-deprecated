package service

import (
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

type ImageService interface {
	CreateImage() (*model.Image, error)
	UpdateImage(id string) (*model.Image, error)
	DeleteImage(id string) (*model.Image, error)
}

type imageService struct{
	database database.Database
}

func NewImageService(db database.Database) ImageService {
	return &imageService{
		database: db,
	}
}

func (s *imageService) CreateImage() (*model.Image, error) {
	panic("not implemented")
}

func (s *imageService) UpdateImage(id string) (*model.Image, error) {
	panic("not implemented")
}

func (s *imageService) DeleteImage(id string) (*model.Image, error) {
	panic("not implemented")
}