package icon

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/common/enum"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"strings"
)

type Repository interface {
	FindAllIcon(*[]*model.Icon) error
	FindOneIcon(int64, *model.Icon) error
	CreateIcon(*model.Icon) error
	UpdateIcon(int64, *model.Icon) error
	DeleteIcon(int64, *model.Icon) error
}

type Service struct {
	repository       Repository
	validatorService service.ValidatorService
}

func NewIconService(repository Repository, validatorService service.ValidatorService) Service {
	return Service{
		repository:       repository,
		validatorService: validatorService,
	}
}

func (s *Service) FindAll() (*[]*model.Icon, error) {
	var icons []*model.Icon

	err := s.repository.FindAllIcon(&icons)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return &icons, nil
}

func (s *Service) FindOne(id int64) (*model.Icon, error) {
	var icon model.Icon

	err := s.repository.FindOneIcon(id, &icon)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &icon, nil
}

func (s *Service) Create(iconDto model.NewIcon) (*model.Icon, error) {
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

func (s *Service) Update(id int64, iconDto model.NewIcon) (*model.Icon, error) {
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

func (s *Service) Delete(id int64) (*model.Icon, error) {
	var icon model.Icon
	err := s.repository.DeleteIcon(id, &icon)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &icon, nil
}

func (s *Service) CheckIconType(iconType enum.IconType) (string, error) {
	err := strings.ToLower(string(iconType))
	if err != string(enum.ICON) && err != string(enum.SVG) {
		return "", fiber.NewError(fiber.StatusBadRequest, "Invalid icon type")
	}
	return err, nil
}

func (s *Service) DtoToRaw(iconDto model.NewIcon) (*model.Icon, error) {
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
