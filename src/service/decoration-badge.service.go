package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
)

type BadgeService interface {
	GetAll() ([]*model.Badge, error)
	GetOne(id int64) (*model.Badge, error)
	Create(badgeDto *model.NewBadge) (*model.Badge, error)
	Update(id int64, badgeDto *model.NewBadge) (*model.Badge, error)
	Delete(id int64) (*model.Badge, error)
	DtoToRaw(githubRepoDto model.NewBadge) (*model.Badge, error)
}

type badgeService struct {
	repository       repository.GormRepository
	iconService      IconService
	validatorService ValidatorService
}

func NewBadgeService(repository repository.GormRepository, iconService IconService, validatorService ValidatorService) BadgeService {
	return &badgeService{
		repository:       repository,
		iconService:      iconService,
		validatorService: validatorService,
	}
}

func (s *badgeService) GetAll() ([]*model.Badge, error) {
	var badge []*model.Badge

	err := s.repository.FindAllBadge(&badge)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return badge, nil
}

func (s *badgeService) GetOne(id int64) (*model.Badge, error) {
	var badge model.Badge

	err := s.repository.FindBadge(id, &badge)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &badge, nil
}

func (s *badgeService) Create(badgeDto *model.NewBadge) (*model.Badge, error) {
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

func (s *badgeService) Update(id int64, badgeDto *model.NewBadge) (*model.Badge, error) {
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

func (s *badgeService) Delete(id int64) (*model.Badge, error) {
	var badge model.Badge
	err := s.repository.DeleteBadge(id, &badge)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &badge, nil
}

func (s badgeService) DtoToRaw(badgeDto model.NewBadge) (*model.Badge, error) {
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
