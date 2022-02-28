package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

type ImageService interface {
	GetAll() ([]*model.Image, error)
	GetOne(id int64) (*model.Image, error)
	Create(imageDto *model.NewImage) (*model.Image, error)
	Update(id int64, imageDto *model.NewImage) (*model.Image, error)
	Delete(id int64) (*model.Image, error)
}

type imageService struct{
	database database.Database
}

func NewImageService(db database.Database) ImageService {
	return &imageService{
		database: db,
	}
}

func (s *imageService) GetAll() ([]*model.Image, error) {
	db := s.database.GetConnection()
	
	var images []*model.Image
	result := db.Find(&images)
	
	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}
	
	return images, nil
}

func (s *imageService) GetOne(id int64) (*model.Image, error) {
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

func (s *imageService) Create(imageDto *model.NewImage) (*model.Image, error) {
	db := s.database.GetConnection()

	image := model.Image{Name:imageDto.Name, Description:imageDto.Description, ImgUrl: imageDto.ImgURL }
	
	result := db.Create(&image)
	
	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &image, nil
}

func (s *imageService) Update(id int64, imageDto *model.NewImage) (*model.Image, error) {
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

func (s *imageService) Delete(id int64) (*model.Image, error) {
	db := s.database.GetConnection()

	var image *model.Image

	result := db.First(&image, id).Delete(&model.Image{}, id)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return image, nil
}