package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

type ImageService interface {
	GetAll() ([]*model.Image, error)
	GetOne(id int64) (*model.Image, error)
	Create(imageDto *model.NewImage) (*model.Image, error)
	Update(id int64, imageDto *model.NewImage) (*model.Image, error)
	Delete(id int64) (*model.Image, error)
	DtoToRaw(imageDto model.NewImage) (*model.Image, error)
}

type imageService struct {
	database         database.Database
	validatorService ValidatorService
}

func NewImageService(db database.Database, validatorService ValidatorService) ImageService {
	return &imageService{
		database:         db,
		validatorService: validatorService,
	}
}

func (s *imageService) GetAll() ([]*model.Image, error) {
	db := s.database.GetConnection()

	var images []*model.Image
	result := db.Find(&images)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return images, nil
}

func (s *imageService) GetOne(id int64) (*model.Image, error) {
	db := s.database.GetConnection()

	var image *model.Image
	result := db.First(&image, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return image, nil
}

func (s *imageService) Create(imageDto *model.NewImage) (*model.Image, error) {
	db := s.database.GetConnection()

	image, err := s.DtoToRaw(*imageDto)
	if err != nil {
		return nil, err
	}

	result := db.Create(&image)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return image, nil
}

func (s *imageService) Update(id int64, imageDto *model.NewImage) (*model.Image, error) {
	db := s.database.GetConnection()

	var image *model.Image
	raw, err := s.DtoToRaw(*imageDto)
	if err != nil {
		return nil, err
	}

	result := db.First(&image, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return image, nil
}

func (s *imageService) Delete(id int64) (*model.Image, error) {
	db := s.database.GetConnection()

	var image *model.Image

	result := db.First(&image, id).Delete(&model.Image{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return image, nil
}

func (s *imageService) DtoToRaw(imageDto model.NewImage) (*model.Image, error) {
	err := s.validatorService.Image(imageDto)
	if err != nil {
		return nil, err
	}

	image := model.Image{
		ID:          imageDto.ID,
		Name:        imageDto.Name,
		Description: imageDto.Description,
		ImgUrl:      imageDto.ImgURL,
	}

	if imageDto.OwnerID > 0 {
		image.OwnerID = imageDto.OwnerID
		image.OwnerType = imageDto.OwnerType
	}

	return &image, nil
}
