package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"gorm.io/gorm"
)

type BadgeService interface {
	GetAll() ([]*model.Badge, error)
	GetOne(id int64) (*model.Badge, error)
	Create(badgeDto *model.NewBadge) (*model.Badge, error)
	Update(id int64, badgeDto *model.NewBadge) (*model.Badge, error)
	Delete(id int64) (*model.Badge, error)
}

type badgeService struct {
	database database.Database
	iconService IconService
}

func NewBadgeService(database database.Database, iconService IconService) BadgeService {
	return &badgeService{
		database: database,
		iconService: iconService,
	}
}

func (s *badgeService) GetAll() ([]*model.Badge, error) {
	db := s.database.GetConnection()

	var badge []*model.Badge

	result := db.Preload("Icon").Find(&badge)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return badge, nil
}

func (s *badgeService) GetOne(id int64) (*model.Badge, error) {
	//TODO: Optimize code to be more efficiency

	db := s.database.GetConnection()

	var badge *model.Badge

	result := db.Preload("Icon").First(&badge, id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return badge, nil
}

func (s *badgeService) Create(badgeDto *model.NewBadge) (*model.Badge, error) {
	icon := model.Icon{Name: badgeDto.Icon.Name, BgColor: badgeDto.Icon.BgColor, IconType: enum.IconType(badgeDto.Icon.IconType)}

	db := s.database.GetConnection()

	badge := model.Badge{Name: badgeDto.Name, Color: badgeDto.Color, Icon: icon}

	result := db.Create(&badge)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &badge, nil
}

func (s *badgeService) Update(id int64, badgeDto *model.NewBadge) (*model.Badge, error) {
	//FIXME: It update despite it not found an icon (Need to make it update correctly)

	db := s.database.GetConnection()

	var badge *model.Badge

	result := db.First(&badge, "id = ?", id).Updates(model.Badge{Name: badgeDto.Name, Color: badgeDto.Color})

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	if badgeDto.Icon.ID != 0 {
		icon, err := s.iconService.GetOne(badgeDto.Icon.ID)
		if err != nil {
			return nil, err
		}

		badge.Icon = *icon
	}
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&badge)

	return badge, nil
}

func (s *badgeService) Delete(id int64) (*model.Badge, error) {
	db := s.database.GetConnection()

	badge, err := s.GetOne(id)

	if err != nil {
		return nil, err
	}

	result := db.Delete(&model.Badge{}, id)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return badge, nil
}
