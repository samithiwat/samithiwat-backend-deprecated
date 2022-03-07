package image

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/service"
)

type Repository interface {
	FindAllImage(*[]*model.Image) error
	FindOneImage(int64, *model.Image) error
	CreateImage(*model.Image) error
	UpdateImage(int64, *model.Image) error
	DeleteImage(int64, *model.Image) error
}

type Service struct {
	repository       Repository
	validatorService service.ValidatorService
}

func NewImageService(repository Repository, validatorService service.ValidatorService) Service {
	return Service{
		repository:       repository,
		validatorService: validatorService,
	}
}

func (s *Service) GetAll() ([]*model.Image, error) {
	var images []*model.Image
	err := s.repository.FindAllImage(&images)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return images, nil
}

func (s *Service) GetOne(id int64) (*model.Image, error) {
	var image model.Image
	err := s.repository.FindOneImage(id, &image)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &image, nil
}

func (s *Service) Create(imageDto *model.NewImage) (*model.Image, error) {
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

func (s *Service) Update(id int64, imageDto *model.NewImage) (*model.Image, error) {
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

func (s *Service) Delete(id int64) (*model.Image, error) {
	var image model.Image
	err := s.repository.DeleteImage(id, &image)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &image, nil
}

func (s *Service) DtoToRaw(imageDto model.NewImage) (*model.Image, error) {
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
