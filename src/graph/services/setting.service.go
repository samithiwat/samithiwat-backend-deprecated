package service

import (
	"github.com/gofiber/fiber/v2"
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
	//TODO: Complete this

	db := s.database.GetConnection()

	setting := model.Setting{AboutMe: settingDto.AboutMe, Timeline: settingDto.Timeline}

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &setting, nil
}

func (s *settingService) Update(id int64, settingDto *model.NewSetting) (*model.Setting, error) {
	//TODO: Complete this

	//db := s.database.GetConnection()

	//setting := model.Setting{AboutMe: model.AboutMe{}}

	//result := db.Create(&setting)
	//
	//if result.Error != nil {
	//	return nil, fiber.ErrUnprocessableEntity
	//}
	//
	//return &setting, nil
	return nil, nil
}

func (s *settingService) Delete(id int64) (*model.Setting, error) {
	//TODO: Complete this

	//db := s.database.GetConnection()

	//setting := model.Setting{AboutMe: model.AboutMe{}}

	//result := db.Create(&setting)
	//
	//if result.Error != nil {
	//	return nil, fiber.ErrUnprocessableEntity
	//}
	//
	//return &setting, nil
	return nil, nil
}
