package setting

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/model"
	repository "github.com/samithiwat/samithiwat-backend/src/repository/gorm"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/service/aboutme"
	"github.com/samithiwat/samithiwat-backend/src/service/timeline"
)

type Service interface {
	GetAll() ([]*model.Setting, error)
	GetOne(id int64) (*model.Setting, error)
	GetActivatedSetting() (*model.Setting, error)
	Create(settingDto *model.NewSetting) (*model.Setting, error)
	Update(id int64, imageDto *model.NewSetting) (*model.Setting, error)
	Delete(id int64) (*model.Setting, error)
	DtoToRaw(settingDto *model.NewSetting) (*model.Setting, error)
}

func NewSettingService(repository repository.GormRepository, aboutMeSettingService aboutme.Service, timelineSettingService timeline.Service, validatorService service.ValidatorService) Service {
	return &settingService{
		repository:             repository,
		aboutMeSettingService:  aboutMeSettingService,
		timelineSettingService: timelineSettingService,
		validatorService:       validatorService,
	}
}

type settingService struct {
	repository             repository.GormRepository
	aboutMeSettingService  aboutme.Service
	timelineSettingService timeline.Service
	validatorService       service.ValidatorService
}

func (s *settingService) GetAll() ([]*model.Setting, error) {
	var settings []*model.Setting

	err := s.repository.FindAllSetting(&settings)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return settings, nil
}

func (s *settingService) GetOne(id int64) (*model.Setting, error) {
	var setting model.Setting

	err := s.repository.FindSetting(id, &setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *settingService) GetActivatedSetting() (*model.Setting, error) {
	var setting model.Setting

	err := s.repository.FindActiveSetting(&setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *settingService) Create(settingDto *model.NewSetting) (*model.Setting, error) {
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

func (s *settingService) Update(id int64, settingDto *model.NewSetting) (*model.Setting, error) {
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

func (s *settingService) Delete(id int64) (*model.Setting, error) {
	var setting model.Setting
	err := s.repository.DeleteSetting(id, &setting)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return &setting, nil
}

func (s *settingService) DtoToRaw(settingDto *model.NewSetting) (*model.Setting, error) {
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
