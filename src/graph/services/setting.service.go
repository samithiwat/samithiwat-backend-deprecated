package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

type SettingService interface {
	GetAll() ([]*model.Setting, error)
	GetOne(id int64) (*model.Setting, error)
	GetActivatedSetting() (*model.Setting, error)
	Create(settingDto *model.NewSetting) (*model.Setting, error)
	Update(id int64, imageDto *model.NewSetting) (*model.Setting, error)
	Delete(id int64) (*model.Setting, error)
	DtoToRaw(settingDto *model.NewSetting) (*model.Setting, error)
}

func NewSettingService(db database.Database, aboutMeSettingService AboutMeSettingService, timelineSettingService TimelineSettingService, validatorService ValidatorService) SettingService {
	return &settingService{
		database:               db,
		aboutMeSettingService:  aboutMeSettingService,
		timelineSettingService: timelineSettingService,
		validatorService: validatorService,
	}
}

type settingService struct {
	database               database.Database
	aboutMeSettingService  AboutMeSettingService
	timelineSettingService TimelineSettingService
	validatorService ValidatorService
}

func (s *settingService) GetAll() ([]*model.Setting, error) {
	db := s.database.GetConnection()

	var settings []*model.Setting

	result := db.Preload("AboutMe").Preload("Timeline").Find(&settings)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return settings, nil
}

func (s *settingService) GetOne(id int64) (*model.Setting, error) {
	db := s.database.GetConnection()

	var setting *model.Setting

	result := db.Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").First(&setting, id)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return setting, nil
}

func (s *settingService) GetActivatedSetting() (*model.Setting, error) {
	db := s.database.GetConnection()

	var setting *model.Setting

	result := db.Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").Where("isActivated = ?", true).Take(&setting)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return setting, nil
}

func (s *settingService) Create(settingDto *model.NewSetting) (*model.Setting, error) {
	db := s.database.GetConnection()
	setting, err := s.DtoToRaw(settingDto)
	if err != nil{
		return nil, err
	}

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return setting, nil
}

func (s *settingService) Update(id int64, settingDto *model.NewSetting) (*model.Setting, error) {
	db := s.database.GetConnection()

	var setting *model.Setting
	raw, err := s.DtoToRaw(settingDto)
	if err != nil{
		return nil, err
	}

	result := db.Preload("AboutMe").Preload("Timeline").Preload("Timeline.Icon").Preload("Timeline.Images").First(&setting, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	if (!cmp.Equal(raw.Timeline, model.Timeline{})) {
		db.Model(&setting).Association("Timeline").Replace(&raw.Timeline)
		db.Model(&setting.Timeline).Association("Icon").Replace(&raw.Timeline.Icon)
		db.Model(&setting.Timeline).Association("Images").Replace(&raw.Timeline.Images)
	}

	if (!cmp.Equal(raw.AboutMe, model.AboutMe{})) {
		db.Model(&setting).Association("AboutMe").Replace(&raw.AboutMe)
	}

	return setting, nil
}

func (s *settingService) Delete(id int64) (*model.Setting, error) {
	db := s.database.GetConnection()

	var setting *model.Setting

	result := db.First(&setting, id).Delete(&model.Setting{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return setting, nil
}

func (s *settingService) DtoToRaw(settingDto *model.NewSetting) (*model.Setting, error) {
	err := s.validatorService.Setting(*settingDto)
	if err != nil{
		return nil, err
	}

	rawTimeline, err := s.timelineSettingService.DtoToRaw(&settingDto.Timeline)
	if err != nil{
		return nil, err
	}
	rawAboutMe, err := s.aboutMeSettingService.DtoToRaw(&settingDto.AboutMe)
	if err != nil{
		return nil, err
	}

	setting := model.Setting{
		AboutMe: *rawAboutMe,
		Timeline: *rawTimeline,
		IsActivated: settingDto.IsActivated,
	}

	return &setting, nil
}
