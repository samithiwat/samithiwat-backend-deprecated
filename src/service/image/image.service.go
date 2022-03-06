package image

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
	"github.com/samithiwat/samithiwat-backend/src/service"
)

type Service interface {
	GetAll() ([]*model.Image, error)
	GetOne(id int64) (*model.Image, error)
	Create(imageDto *model.NewImage) (*model.Image, error)
	Update(id int64, imageDto *model.NewImage) (*model.Image, error)
	Delete(id int64) (*model.Image, error)
	DtoToRaw(imageDto model.NewImage) (*model.Image, error)
}

type imageService struct {
	repository       repository.GormRepository
	validatorService service.ValidatorService
}

func NewImageService(repository repository.GormRepository, validatorService service.ValidatorService) Service {
	return &imageService{
		repository:       repository,
		validatorService: validatorService,
	}
}

func (s *imageService) GetAll() ([]*model.Image, error) {
	var images []*model.Image
	err := s.repository.FindAllImage(&images)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return images, nil
}

func (s *imageService) GetOne(id int64) (*model.Image, error) {
	var image model.Image
	err := s.repository.FindImage(id, &image)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &image, nil
}

func (s *imageService) Create(imageDto *model.NewImage) (*model.Image, error) {
	image, err := s.DtoToRaw(*imageDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateImage(image)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return image, nil
}

func (s *imageService) Update(id int64, imageDto *model.NewImage) (*model.Image, error) {
	image, err := s.DtoToRaw(*imageDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.UpdateImage(id, image)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return image, nil
}

func (s *imageService) Delete(id int64) (*model.Image, error) {
	var image model.Image
	err := s.repository.DeleteImage(id, &image)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &image, nil
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
