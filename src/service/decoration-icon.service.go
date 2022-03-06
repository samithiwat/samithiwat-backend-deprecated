package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"strings"
)

type IconService interface {
	GetAll() ([]*model.Icon, error)
	GetOne(id int64) (*model.Icon, error)
	Create(iconDto model.NewIcon) (*model.Icon, error)
	Update(id int64, iconDto model.NewIcon) (*model.Icon, error)
	Delete(id int64) (*model.Icon, error)
	CheckIconType(iconType enum.IconType) (string, error)
	DtoToRaw(iconDto model.NewIcon) (*model.Icon, error)
}

type iconService struct {
	database         database.Database
	validatorService ValidatorService
}

func NewIconService(database database.Database, validatorService ValidatorService) IconService {
	return &iconService{
		database:         database,
		validatorService: validatorService,
	}
}

func (s *iconService) GetAll() ([]*model.Icon, error) {
	db := s.database.GetConnection()

	var icons []*model.Icon

	result := db.Find(&icons)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return icons, nil
}

func (s *iconService) GetOne(id int64) (*model.Icon, error) {
	db := s.database.GetConnection()

	var icon *model.Icon

	result := db.First(&icon, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return icon, nil
}

func (s *iconService) Create(iconDto model.NewIcon) (*model.Icon, error) {
	db := s.database.GetConnection()

	icon, err := s.DtoToRaw(iconDto)
	if err != nil {
		return nil, err
	}

	result := db.Create(&icon)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return icon, nil
}

func (s *iconService) Update(id int64, iconDto model.NewIcon) (*model.Icon, error) {
	db := s.database.GetConnection()

	var icon *model.Icon
	raw, err := s.DtoToRaw(iconDto)
	if err != nil {
		return nil, err
	}

	result := db.First(&icon, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return icon, nil
}

func (s *iconService) Delete(id int64) (*model.Icon, error) {
	db := s.database.GetConnection()

	var icon *model.Icon

	result := db.First(&icon, id).Delete(&model.Icon{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return icon, nil
}

func (s *iconService) CheckIconType(iconType enum.IconType) (string, error) {
	result := strings.ToLower(string(iconType))
	if result != string(enum.ICON) && result != string(enum.SVG) {
		return "", fiber.NewError(fiber.StatusBadRequest, "Invalid icon type")
	}
	return result, nil
}

func (s *iconService) DtoToRaw(iconDto model.NewIcon) (*model.Icon, error) {
	err := s.validatorService.Icon(iconDto)
	if err != nil {
		return nil, err
	}

	icon := model.Icon{
		ID:       iconDto.ID,
		Name:     iconDto.Name,
		BgColor:  iconDto.BgColor,
		IconType: enum.IconType(iconDto.IconType),
	}

	if iconDto.OwnerID > 0 {
		icon.OwnerID = iconDto.OwnerID
		icon.OwnerType = iconDto.OwnerType
	}

	return &icon, nil
}
