package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	model2 "github.com/samithiwat/samithiwat-backend/src/model"
)

type ImageService interface {
	GetAll() ([]*model2.Image, error)
	GetOne(id int64) (*model2.Image, error)
	Create(imageDto *model2.NewImage) (*model2.Image, error)
	Update(id int64, imageDto *model2.NewImage) (*model2.Image, error)
	Delete(id int64) (*model2.Image, error)
	DtoToRaw(imageDto model2.NewImage) (*model2.Image, error)
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

func (s *imageService) GetAll() ([]*model2.Image, error) {
	db := s.database.GetConnection()

	var images []*model2.Image
	result := db.Find(&images)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return images, nil
}

func (s *imageService) GetOne(id int64) (*model2.Image, error) {
	db := s.database.GetConnection()

	var image *model2.Image
	result := db.First(&image, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return image, nil
}

func (s *imageService) Create(imageDto *model2.NewImage) (*model2.Image, error) {
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

func (s *imageService) Update(id int64, imageDto *model2.NewImage) (*model2.Image, error) {
	db := s.database.GetConnection()

	var image *model2.Image
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

func (s *imageService) Delete(id int64) (*model2.Image, error) {
	db := s.database.GetConnection()

	var image *model2.Image

	result := db.First(&image, id).Delete(&model2.Image{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return image, nil
}

func (s *imageService) DtoToRaw(imageDto model2.NewImage) (*model2.Image, error) {
	err := s.validatorService.Image(imageDto)
	if err != nil {
		return nil, err
	}

	image := model2.Image{
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
