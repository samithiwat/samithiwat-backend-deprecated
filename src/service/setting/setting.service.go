package setting

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/service/aboutme"
	"github.com/samithiwat/samithiwat-backend/src/service/timeline"
)

type Repository interface {
	FindAllSetting(*[]*model.Setting) error
	FindOneSetting(int64, *model.Setting) error
	FindActiveSetting(*model.Setting) error
	CreateSetting(*model.Setting) error
	UpdateSetting(int64, *model.Setting) error
	DeleteSetting(int64, *model.Setting) error
}

type Service struct {
	repository             Repository
	validatorService       service.ValidatorService
	aboutMeSettingService  aboutme.Service
	timelineSettingService timeline.Service
}

func NewSettingService(repository Repository, aboutMeSettingService aboutme.Service, timelineSettingService timeline.Service, validatorService service.ValidatorService) Service {
	return Service{
		repository:             repository,
		aboutMeSettingService:  aboutMeSettingService,
		timelineSettingService: timelineSettingService,
		validatorService:       validatorService,
	}
}

func (s *Service) GetAll() ([]*model.Setting, error) {
	var settings []*model.Setting

	err := s.repository.FindAllSetting(&settings)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return settings, nil
}

func (s *Service) GetOne(id int64) (*model.Setting, error) {
	var setting model.Setting

	err := s.repository.FindOneSetting(id, &setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *Service) GetActivatedSetting() (*model.Setting, error) {
	var setting model.Setting

	err := s.repository.FindActiveSetting(&setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *Service) Create(settingDto *model.NewSetting) (*model.Setting, error) {
	setting, err := s.DtoToRaw(settingDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateSetting(setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return setting, nil
}

func (s *Service) Update(id int64, settingDto *model.NewSetting) (*model.Setting, error) {
	setting, err := s.DtoToRaw(settingDto)
	if err != nil {
		return nil, err
	}

	err = s.repository.UpdateSetting(id, setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return setting, nil
}

func (s *Service) Delete(id int64) (*model.Setting, error) {
	var setting model.Setting
	err := s.repository.DeleteSetting(id, &setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *Service) DtoToRaw(settingDto *model.NewSetting) (*model.Setting, error) {
	err := s.validatorService.Setting(*settingDto)
	if err != nil {
		return nil, err
	}

	rawTimeline, err := s.timelineSettingService.DtoToRaw(&settingDto.Timeline)
	if err != nil {
		return nil, err
	}
	rawAboutMe, err := s.aboutMeSettingService.DtoToRaw(&settingDto.AboutMe)
	if err != nil {
		return nil, err
	}

	setting := model.Setting{
		AboutMe:     *rawAboutMe,
		Timeline:    *rawTimeline,
		IsActivated: settingDto.IsActivated,
	}

	return &setting, nil
}
