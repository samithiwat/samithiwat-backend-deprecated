package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/database"
	model2 "github.com/samithiwat/samithiwat-backend/src/model"
)

type SettingService interface {
	GetAll() ([]*model2.Setting, error)
	GetOne(id int64) (*model2.Setting, error)
	GetActivatedSetting() (*model2.Setting, error)
	Create(settingDto *model2.NewSetting) (*model2.Setting, error)
	Update(id int64, imageDto *model2.NewSetting) (*model2.Setting, error)
	Delete(id int64) (*model2.Setting, error)
	DtoToRaw(settingDto *model2.NewSetting) (*model2.Setting, error)
}

func NewSettingService(db database.Database, aboutMeSettingService AboutMeSettingService, timelineSettingService TimelineSettingService, validatorService ValidatorService) SettingService {
	return &settingService{
		database:               db,
		aboutMeSettingService:  aboutMeSettingService,
		timelineSettingService: timelineSettingService,
		validatorService:       validatorService,
	}
}

type settingService struct {
	database               database.Database
	aboutMeSettingService  AboutMeSettingService
	timelineSettingService TimelineSettingService
	validatorService       ValidatorService
}

func (s *settingService) GetAll() ([]*model2.Setting, error) {
	db := s.database.GetConnection()

	var settings []*model2.Setting

	result := db.Preload("AboutMe").Preload("Timeline").Find(&settings)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return settings, nil
}

func (s *settingService) GetOne(id int64) (*model2.Setting, error) {
	db := s.database.GetConnection()

	var setting *model2.Setting

	result := db.Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").First(&setting, id)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return setting, nil
}

func (s *settingService) GetActivatedSetting() (*model2.Setting, error) {
	db := s.database.GetConnection()

	var setting *model2.Setting

	result := db.Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").Where("isActivated = ?", true).Take(&setting)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return setting, nil
}

func (s *settingService) Create(settingDto *model2.NewSetting) (*model2.Setting, error) {
	db := s.database.GetConnection()
	setting, err := s.DtoToRaw(settingDto)
	if err != nil {
		return nil, err
	}

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return setting, nil
}

func (s *settingService) Update(id int64, settingDto *model2.NewSetting) (*model2.Setting, error) {
	db := s.database.GetConnection()

	var setting *model2.Setting
	raw, err := s.DtoToRaw(settingDto)
	if err != nil {
		return nil, err
	}

	result := db.Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").First(&setting, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	if (!cmp.Equal(raw.Timeline, model2.Timeline{})) {
		db.Model(&setting).Association("Timeline").Replace(&raw.Timeline)
		db.Model(&setting.Timeline).Association("Icon").Replace(&raw.Timeline.Icon)
		db.Model(&setting.Timeline).Association("Images").Replace(&raw.Timeline.Images)
	}

	if (!cmp.Equal(raw.AboutMe, model2.AboutMe{})) {
		db.Model(&setting).Association("AboutMe").Replace(&raw.AboutMe)
	}

	return setting, nil
}

func (s *settingService) Delete(id int64) (*model2.Setting, error) {
	db := s.database.GetConnection()

	var setting *model2.Setting

	result := db.First(&setting, id).Delete(&model2.Setting{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return setting, nil
}

func (s *settingService) DtoToRaw(settingDto *model2.NewSetting) (*model2.Setting, error) {
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

	setting := model2.Setting{
		AboutMe:     *rawAboutMe,
		Timeline:    *rawTimeline,
		IsActivated: settingDto.IsActivated,
	}

	return &setting, nil
}
