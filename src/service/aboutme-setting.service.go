package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
)

type AboutMeSettingService interface {
	GetAll() ([]*model.AboutMe, error)
	GetOne(id int64) (*model.AboutMe, error)
	Create(settingDto *model.NewAboutMe) (*model.AboutMe, error)
	Update(id int64, settingDto *model.NewAboutMe) (*model.AboutMe, error)
	Delete(id int64) (*model.AboutMe, error)
	DtoToRaw(settingDto *model.NewAboutMe) (*model.AboutMe, error)
}

func NewAboutMeSettingService(repository repository.GormRepository, validatorService ValidatorService) AboutMeSettingService {
	return &aboutMeSettingService{
		repository:       repository,
		validatorService: validatorService,
	}
}

type aboutMeSettingService struct {
	repository       repository.GormRepository
	validatorService ValidatorService
}

func (s *aboutMeSettingService) GetAll() ([]*model.AboutMe, error) {
	var settings []*model.AboutMe

	err := s.repository.FindAllAboutMe(&settings)

	if err.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return settings, nil
}

func (s *aboutMeSettingService) GetOne(id int64) (*model.AboutMe, error) {
	var setting model.AboutMe

	err := s.repository.FindAboutMe(id, &setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *aboutMeSettingService) Create(aboutMeDto *model.NewAboutMe) (*model.AboutMe, error) {
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

func (s *aboutMeSettingService) Update(id int64, aboutMeDto *model.NewAboutMe) (*model.AboutMe, error) {
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

func (s *aboutMeSettingService) Delete(id int64) (*model.AboutMe, error) {
	var setting model.AboutMe
	err := s.repository.DeleteAboutMe(id, &setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *aboutMeSettingService) DtoToRaw(settingDto *model.NewAboutMe) (*model.AboutMe, error) {
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
