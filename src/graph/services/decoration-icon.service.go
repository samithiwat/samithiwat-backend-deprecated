package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"strings"
)

type IconService interface {
	GetAll() ([]*model.Icon, error)
	GetOne(id int64) (*model.Icon, error)
	Create(iconDto model.NewIcon) (*model.Icon, error)
	Update(id int64, iconDto model.NewIcon) (*model.Icon, error)
	Delete(id int64) (*model.Icon, error)
	CheckIconType(iconType enum.IconType) (string, error)
}

type iconService struct {
	database database.Database
}

func NewIconService(database database.Database) IconService {
	return &iconService{
		database: database,
	}
}

func (s *iconService) GetAll() ([]*model.Icon, error) {
	db := s.database.GetConnection()

	var icons []*model.Icon

	result := db.Find(&icons)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return icons, nil
}

func (s *iconService) GetOne(id int64) (*model.Icon, error) {
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

func (s *iconService) Create(iconDto model.NewIcon) (*model.Icon, error) {
	db := s.database.GetConnection()

	icon := model.Icon{Name: iconDto.Name, BgColor: iconDto.BgColor, IconType: enum.IconType(iconDto.IconType)}

	result := db.Create(&icon)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &icon, nil
}

func (s *iconService) Update(id int64, iconDto model.NewIcon) (*model.Icon, error) {
	db := s.database.GetConnection()

	var icon *model.Icon

	result := db.First(&icon, "id = ?", id).Updates(model.Icon{Name: iconDto.Name, BgColor: iconDto.BgColor, IconType: enum.IconType(iconDto.IconType)})

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return icon, nil
}

func (s *iconService) Delete(id int64) (*model.Icon, error) {
	db := s.database.GetConnection()

	icon, err := s.GetOne(id)

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

func (s *iconService) CheckIconType(iconType enum.IconType) (string, error) {
	result := strings.ToLower(string(iconType))
	if result != "icon" && result != "svg" {
		return "", fiber.ErrBadRequest
	}
	return result, nil
}