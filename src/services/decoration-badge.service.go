package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/database"
	model2 "github.com/samithiwat/samithiwat-backend/src/model"
)

type BadgeService interface {
	GetAll() ([]*model2.Badge, error)
	GetOne(id int64) (*model2.Badge, error)
	Create(badgeDto *model2.NewBadge) (*model2.Badge, error)
	Update(id int64, badgeDto *model2.NewBadge) (*model2.Badge, error)
	Delete(id int64) (*model2.Badge, error)
	DtoToRaw(githubRepoDto model2.NewBadge) (*model2.Badge, error)
}

type badgeService struct {
	database         database.Database
	iconService      IconService
	validatorService ValidatorService
}

func NewBadgeService(database database.Database, iconService IconService, validatorService ValidatorService) BadgeService {
	return &badgeService{
		database:         database,
		iconService:      iconService,
		validatorService: validatorService,
	}
}

func (s *badgeService) GetAll() ([]*model2.Badge, error) {
	db := s.database.GetConnection()

	var badge []*model2.Badge

	result := db.Find(&badge)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return badge, nil
}

func (s *badgeService) GetOne(id int64) (*model2.Badge, error) {
	db := s.database.GetConnection()

	var badge *model2.Badge

	result := db.Preload("Icon").First(&badge, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return badge, nil
}

func (s *badgeService) Create(badgeDto *model2.NewBadge) (*model2.Badge, error) {
	db := s.database.GetConnection()

	badge, err := s.DtoToRaw(*badgeDto)
	if err != nil {
		return nil, err
	}

	result := db.Create(&badge)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return badge, nil
}

func (s *badgeService) Update(id int64, badgeDto *model2.NewBadge) (*model2.Badge, error) {
	db := s.database.GetConnection()

	var badge *model2.Badge
	raw, err := s.DtoToRaw(*badgeDto)
	if err != nil {
		return nil, err
	}

	result := db.Preload("Icon").First(&badge, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	if (!cmp.Equal(raw.Icon, model2.Icon{})) {
		db.Model(&badge).Association("Icon").Replace(&raw.Icon)
	}

	return badge, nil
}

func (s *badgeService) Delete(id int64) (*model2.Badge, error) {
	db := s.database.GetConnection()

	var badge *model2.Badge

	result := db.Preload("Icon").First(&badge, id).Delete(&model2.Badge{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return badge, nil
}

func (s badgeService) DtoToRaw(badgeDto model2.NewBadge) (*model2.Badge, error) {
	err := s.validatorService.Badge(badgeDto)
	if err != nil {
		return nil, err
	}

	rawIcon, err := s.iconService.DtoToRaw(badgeDto.Icon)
	if err != nil {
		return nil, err
	}
	badge := model2.Badge{
		ID:    badgeDto.ID,
		Name:  badgeDto.Name,
		Color: badgeDto.Color,
		Icon:  *rawIcon,
	}

	return &badge, nil
}
