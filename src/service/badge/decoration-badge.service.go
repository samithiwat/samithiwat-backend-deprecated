package badge

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/service/icon"
)

type Repository interface {
	FindAllBadge(*[]*model.Badge) error
	FindOneBadge(int64, *model.Badge) error
	CreateBadge(*model.Badge) error
	UpdateBadge(int64, *model.Badge) error
	DeleteBadge(int64, *model.Badge) error
}

type Service struct {
	repository       Repository
	iconService      icon.Service
	validatorService service.ValidatorService
}

func NewBadgeService(repository Repository, iconService icon.Service, validatorService service.ValidatorService) Service {
	return Service{
		repository:       repository,
		iconService:      iconService,
		validatorService: validatorService,
	}
}

func (s *Service) GetAll() ([]*model.Badge, error) {
	var badge []*model.Badge

	err := s.repository.FindAllBadge(&badge)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return badge, nil
}

func (s *Service) GetOne(id int64) (*model.Badge, error) {
	var badge model.Badge

	err := s.repository.FindOneBadge(id, &badge)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &badge, nil
}

func (s *Service) Create(badgeDto *model.NewBadge) (*model.Badge, error) {
	badge, err := s.DtoToRaw(*badgeDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateBadge(badge)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return badge, nil
}

func (s *Service) Update(id int64, badgeDto *model.NewBadge) (*model.Badge, error) {
	badge, err := s.DtoToRaw(*badgeDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.UpdateBadge(id, badge)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return badge, nil
}

func (s *Service) Delete(id int64) (*model.Badge, error) {
	var badge model.Badge
	err := s.repository.DeleteBadge(id, &badge)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &badge, nil
}

func (s Service) DtoToRaw(badgeDto model.NewBadge) (*model.Badge, error) {
	err := s.validatorService.Badge(badgeDto)
	if err != nil {
		return nil, err
	}

	rawIcon, err := s.iconService.DtoToRaw(badgeDto.Icon)
	if err != nil {
		return nil, err
	}
	badge := model.Badge{
		ID:    badgeDto.ID,
		Name:  badgeDto.Name,
		Color: badgeDto.Color,
		Icon:  *rawIcon,
	}

	return &badge, nil
}
