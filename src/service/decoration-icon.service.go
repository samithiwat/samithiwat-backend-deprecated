package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"github.com/samithiwat/samithiwat-backend/src/model"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
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
	repository       repository.GormRepository
	validatorService ValidatorService
}

func NewIconService(repository repository.GormRepository, validatorService ValidatorService) IconService {
	return &iconService{
		repository:       repository,
		validatorService: validatorService,
	}
}

func (s *iconService) GetAll() ([]*model.Icon, error) {
	var icons []*model.Icon

	err := s.repository.FindAllIcon(&icons)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return icons, nil
}

func (s *iconService) GetOne(id int64) (*model.Icon, error) {
	var icon model.Icon

	err := s.repository.FindIcon(id, &icon)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &icon, nil
}

func (s *iconService) Create(iconDto model.NewIcon) (*model.Icon, error) {
	icon, err := s.DtoToRaw(iconDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateIcon(icon)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return icon, nil
}

func (s *iconService) Update(id int64, iconDto model.NewIcon) (*model.Icon, error) {
	icon, err := s.DtoToRaw(iconDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.UpdateIcon(id, icon)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return icon, nil
}

func (s *iconService) Delete(id int64) (*model.Icon, error) {
	var icon model.Icon
	err := s.repository.DeleteIcon(id, &icon)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &icon, nil
}

func (s *iconService) CheckIconType(iconType enum.IconType) (string, error) {
	err := strings.ToLower(string(iconType))
	if err != string(enum.ICON) && err != string(enum.SVG) {
		return "", fiber.NewError(fiber.StatusBadRequest, "Invalid icon type")
	}
	return err, nil
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
