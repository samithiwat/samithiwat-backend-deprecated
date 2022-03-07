package aboutme

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/service"
)

type Repository interface {
	FindAllAboutMe(*[]*model.AboutMe) error
	FindOneAboutMe(int64, *model.AboutMe) error
	CreateAboutMe(*model.AboutMe) error
	UpdateAboutMe(int64, *model.AboutMe) error
	DeleteAboutMe(int64, *model.AboutMe) error
}

type Service struct {
	repository       Repository
	validatorService service.ValidatorService
}

func NewAboutMeSettingService(repository Repository, validatorService service.ValidatorService) Service {
	return Service{
		repository:       repository,
		validatorService: validatorService,
	}
}

func (s *Service) GetAll() ([]*model.AboutMe, error) {
	var settings []*model.AboutMe

	err := s.repository.FindAllAboutMe(&settings)

	if err.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return settings, nil
}

func (s *Service) GetOne(id int64) (*model.AboutMe, error) {
	var setting model.AboutMe

	err := s.repository.FindOneAboutMe(id, &setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *Service) Create(aboutMeDto *model.NewAboutMe) (*model.AboutMe, error) {
	setting, err := s.DtoToRaw(aboutMeDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateAboutMe(setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return setting, nil
}

func (s *Service) Update(id int64, aboutMeDto *model.NewAboutMe) (*model.AboutMe, error) {
	var aboutMe model.AboutMe
	raw, err := s.DtoToRaw(aboutMeDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.UpdateAboutMe(id, raw)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &aboutMe, nil
}

func (s *Service) Delete(id int64) (*model.AboutMe, error) {
	var setting model.AboutMe
	err := s.repository.DeleteAboutMe(id, &setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *Service) DtoToRaw(settingDto *model.NewAboutMe) (*model.AboutMe, error) {
	err := s.validatorService.AboutMe(*settingDto)
	if err != nil {
		return nil, err
	}

	aboutMe := model.AboutMe{
		ID:          settingDto.ID,
		Name:        settingDto.Name,
		Description: settingDto.Description,
		Content:     settingDto.Content,
		ImgUrl:      settingDto.ImgURL,
		SettingID:   settingDto.SettingID,
	}

	return &aboutMe, err
}
