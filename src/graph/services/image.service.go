package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

type ImageService interface {
	GetAllImages() ([]*model.Image, error)
	GetImage(id int64) (*model.Image, error)
	CreateImage(imageDto *model.NewImage) (*model.Image, error)
	UpdateImage(id int64, imageDto *model.NewImage) (*model.Image, error)
	DeleteImage(id int64) (*model.Image, error)
}

type imageService struct{
	database database.Database
}

func NewImageService(db database.Database) ImageService {
	return &imageService{
		database: db,
	}
}

func (s *imageService) GetAllImages() ([]*model.Image, error) {
	db := s.database.GetConnection()
	
	var images []*model.Image
	result := db.Find(&images)
	
	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}
	
	return images, nil
}

func (s *imageService) GetImage(id int64) (*model.Image, error) {
	db := s.database.GetConnection()

	
	var image *model.Image
	result := db.First(&image, id)
	
	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}
	
	return image, nil
}

func (s *imageService) CreateImage(imageDto *model.NewImage) (*model.Image, error) {
	db := s.database.GetConnection()

	image := model.Image{Name:imageDto.Name, Description:imageDto.Description, ImgUrl: imageDto.ImgURL }
	
	result := db.Create(&image)
	
	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &image, nil
}

func (s *imageService) UpdateImage(id int64, imageDto *model.NewImage) (*model.Image, error) {
	db := s.database.GetConnection()
	
	image := model.Image{Name:imageDto.Name, Description:imageDto.Description, ImgUrl: imageDto.ImgURL }

	result := db.Model(model.Image{}).Where("id = ?", id).Updates(&image)
	
	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}
	
	return &image, nil
}

func (s *imageService) DeleteImage(id int64) (*model.Image, error) {
	db := s.database.GetConnection()

	image, err := s.GetImage(id)

	if err != nil {
		return nil, err
	}

	result := db.Delete(&model.Image{}, id)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return image, nil
}