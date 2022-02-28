package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
)

type AboutMeSettingService interface {
	GetAll() ([]*model.AboutMe, error)
	GetOne(id int64) (*model.AboutMe, error)
	Create(settingDto *model.NewAboutMe) (*model.AboutMe, error)
	Update(id int64, imageDto *model.NewAboutMe) (*model.AboutMe, error)
	Delete(id int64) (*model.AboutMe, error)
}

func NewAboutMeSettingService(db database.Database) AboutMeSettingService {
	return &aboutMeSettingService{
		database: db,
	}
}

type aboutMeSettingService struct {
	database database.Database
}

func (s *aboutMeSettingService) GetAll() ([]*model.AboutMe, error) {
	db := s.database.GetConnection()

	var settings []*model.AboutMe

	result := db.Find(&settings)
	if result.Error != nil {
		return nil, result.Error
	}

	return settings, nil
}

func (s *aboutMeSettingService) GetOne(id int64) (*model.AboutMe, error) {
	db := s.database.GetConnection()

	var setting *model.AboutMe

	result := db.First(&setting, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return setting, nil
}

func (s *aboutMeSettingService) Create(aboutMeDto *model.NewAboutMe) (*model.AboutMe, error) {
	db := s.database.GetConnection()

	setting := model.AboutMe{Name: aboutMeDto.Name, Description: aboutMeDto.Description, Content: aboutMeDto.Content, ImgUrl: aboutMeDto.ImgURL}

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	return &setting, nil
}

func (s *aboutMeSettingService) Update(id int64, aboutMeDto *model.NewAboutMe) (*model.AboutMe, error) {
	db := s.database.GetConnection()

	var aboutMe *model.AboutMe

	result := db.First(&aboutMe, "id = ?", id).Updates(model.AboutMe{Name: aboutMeDto.Name, Description: aboutMeDto.Description, Content: aboutMeDto.Content, ImgUrl: aboutMeDto.ImgURL})

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return aboutMe, nil
}

func (s *aboutMeSettingService) Delete(id int64) (*model.AboutMe, error) {
	db := s.database.GetConnection()

	var aboutMe *model.AboutMe

	result := db.First(&aboutMe, id).Delete(&model.AboutMe{}, id)

	if result.Error != nil {
		return nil, fiber.ErrUnprocessableEntity
	}

	if result.RowsAffected == 0 {
		return nil, fiber.ErrNotFound
	}

	return aboutMe, nil
}