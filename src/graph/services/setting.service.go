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
	DtoToRaw(settingDto *model.NewSetting) *model.Setting
}

func NewSettingService(db database.Database, aboutMeSettingService AboutMeSettingService, timelineSettingService TimelineSettingService) SettingService {
	return &settingService{
		database:               db,
		aboutMeSettingService:  aboutMeSettingService,
		timelineSettingService: timelineSettingService,
	}
}

type settingService struct {
	database               database.Database
	aboutMeSettingService  AboutMeSettingService
	timelineSettingService TimelineSettingService
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
	setting := s.DtoToRaw(settingDto)

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return setting, nil
}

func (s *settingService) Update(id int64, settingDto *model.NewSetting) (*model.Setting, error) {
	db := s.database.GetConnection()

	var setting *model.Setting
	raw := s.DtoToRaw(settingDto)

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

func (s *settingService) DtoToRaw(settingDto *model.NewSetting) *model.Setting {
	rawTimeline := s.timelineSettingService.DtoToRaw(&settingDto.Timeline)
	rawAboutMe := s.aboutMeSettingService.DtoToRaw(&settingDto.AboutMe)
	setting := model.Setting{AboutMe: *rawAboutMe, Timeline: *rawTimeline, IsActivated: settingDto.IsActivated}

	return &setting
}
