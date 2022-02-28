package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	"gorm.io/gorm"
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

	result := db.Find(&settings)
	if result.Error != nil {
		return nil, result.Error
	}

	return settings, nil
}

func (s *settingService) GetOne(id int64) (*model.Setting, error) {
	db := s.database.GetConnection()

	var setting *model.Setting

	result := db.First(&setting, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return setting, nil
}

func (s *settingService) GetActivatedSetting() (*model.Setting, error) {
	db := s.database.GetConnection()

	var setting *model.Setting

	result := db.Where("isActivated = ?", true).Take(&setting)
	if result.Error != nil {
		return nil, result.Error
	}

	return setting, nil
}

func (s *settingService) Create(settingDto *model.NewSetting) (*model.Setting, error) {
	db := s.database.GetConnection()
	setting := s.DtoToRaw(settingDto)

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return setting, nil
}

func (s *settingService) Update(id int64, settingDto *model.NewSetting) (*model.Setting, error) {
	//TODO: Complete this

	db := s.database.GetConnection()

	var setting *model.Setting
	rawSetting := s.DtoToRaw(settingDto)

	result := db.Omit("SettingID").First(&setting, "id = ?", id).Updates(rawSetting)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&setting)
	return setting, nil
}

func (s *settingService) Delete(id int64) (*model.Setting, error) {
	db := s.database.GetConnection()

	var setting *model.Setting

	result := db.First(&setting, id).Delete(&model.Setting{}, id)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return setting, nil
}

func (s *settingService) DtoToRaw(settingDto *model.NewSetting) *model.Setting {
	rawTimeline := s.timelineSettingService.DtoToRaw(&settingDto.Timeline)
	rawAboutMe := s.aboutMeSettingService.DtoToRaw(&settingDto.AboutMe)
	setting := model.Setting{AboutMe: *rawAboutMe, Timeline: *rawTimeline}

	return &setting
}
