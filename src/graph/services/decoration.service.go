package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"gorm.io/gorm"
	"strings"
)

type DecorationService interface {
	GetAllIcons() ([]*model.Icon, error)
	GetIcon(id int64) (*model.Icon, error)
	CreateIcon(iconDto model.NewIcon) (*model.Icon, error)
	UpdateIcon(id int64, iconDto model.NewIcon) (*model.Icon, error)
	DeleteIcon(id int64) (*model.Icon, error)
	CheckIconType(iconType enum.IconType) (string, error)
	GetAllBadges() ([]*model.Badge, error)
	GetBadge(id int64) (*model.Badge, error)
	CreateBadge(badgeDto *model.NewBadge) (*model.Badge, error)
	UpdateBadge(id int64, badgeDto *model.NewBadge) (*model.Badge, error)
	DeleteBadge(id int64) (*model.Badge, error)
}

type decorationService struct {
	database database.Database
}

func NewDecorationService(database database.Database) DecorationService {
	return &decorationService{
		database: database,
	}
}

func (s *decorationService) GetAllIcons() ([]*model.Icon, error) {
	db := s.database.GetConnection()

	var icons []*model.Icon

	result := db.Find(&icons)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return icons, nil
}

func (s *decorationService) GetIcon(id int64) (*model.Icon, error) {
	db := s.database.GetConnection()

	var icon *model.Icon

	result := db.First(&icon, id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return icon, nil
}

func (s *decorationService) CreateIcon(iconDto model.NewIcon) (*model.Icon, error) {
	db := s.database.GetConnection()

	icon := model.Icon{Name: iconDto.Name, BgColor: iconDto.BgColor, IconType: enum.IconType(iconDto.IconType)}

	result := db.Create(&icon)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &icon, nil
}

func (s *decorationService) UpdateIcon(id int64, iconDto model.NewIcon) (*model.Icon, error) {
	db := s.database.GetConnection()

	icon := model.Icon{Name: iconDto.Name, BgColor: iconDto.BgColor, IconType: enum.IconType(iconDto.IconType)}

	result := db.Model(model.Icon{}).Where("id = ?", id).Updates(&icon)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	iconResult, err := s.GetIcon(id)

	if err != nil {
		return nil, err
	}

	return iconResult, nil
}

func (s *decorationService) DeleteIcon(id int64) (*model.Icon, error) {
	db := s.database.GetConnection()

	icon, err := s.GetIcon(id)

	if err != nil {
		return nil, err
	}

	result := db.Delete(&model.Icon{}, id)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return icon, nil
}

func (s *decorationService) CheckIconType(iconType enum.IconType) (string, error) {
	result := strings.ToLower(string(iconType))
	if result != "icon" && result != "svg" {
		return "", fiber.ErrBadRequest
	}
	return result, nil
}

func (s *decorationService) GetAllBadges() ([]*model.Badge, error) {
	db := s.database.GetConnection()

	var badge []*model.Badge

	result := db.Find(&badge)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return badge, nil
}

func (s *decorationService) GetBadge(id int64) (*model.Badge, error) {
	db := s.database.GetConnection()

	var badge *model.Badge

	result := db.First(&badge, id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return badge, nil
}

func (s *decorationService) CreateBadge(badgeDto *model.NewBadge) (*model.Badge, error) {
	icon := model.Icon{Name: badgeDto.Icon.Name, BgColor: badgeDto.Icon.BgColor, IconType: enum.IconType(badgeDto.Icon.IconType)}

	db := s.database.GetConnection()

	badge := model.Badge{Name: badgeDto.Name, Color: badgeDto.Color, Icon: icon}

	result := db.Create(&badge)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &badge, nil
}

func (s *decorationService) UpdateBadge(id int64, badgeDto *model.NewBadge) (*model.Badge, error) {
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
		icon, err := s.GetIcon(badgeDto.Icon.ID)
		if err != nil {
			return nil, err
		}

		badge.Icon = *icon
	}
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&badge)

	return badge, nil
}

func (s *decorationService) DeleteBadge(id int64) (*model.Badge, error) {
	db := s.database.GetConnection()

	badge, err := s.GetBadge(id)

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
