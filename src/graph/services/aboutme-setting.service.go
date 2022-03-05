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
	DtoToRaw(settingDto *model.NewAboutMe) *model.AboutMe
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
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return settings, nil
}

func (s *aboutMeSettingService) GetOne(id int64) (*model.AboutMe, error) {
	db := s.database.GetConnection()

	var setting *model.AboutMe

	result := db.First(&setting, id)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	return setting, nil
}

func (s *aboutMeSettingService) Create(aboutMeDto *model.NewAboutMe) (*model.AboutMe, error) {
	db := s.database.GetConnection()

	setting := s.DtoToRaw(aboutMeDto)

	result := db.Create(&setting)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return setting, nil
}

func (s *aboutMeSettingService) Update(id int64, aboutMeDto *model.NewAboutMe) (*model.AboutMe, error) {
	db := s.database.GetConnection()

	var aboutMe *model.AboutMe
	raw := s.DtoToRaw(aboutMeDto)

	result := db.First(&aboutMe, "id = ?", id).Updates(raw)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return aboutMe, nil
}

func (s *aboutMeSettingService) Delete(id int64) (*model.AboutMe, error) {
	db := s.database.GetConnection()

	var aboutMe *model.AboutMe

	result := db.First(&aboutMe, id).Delete(&model.AboutMe{}, id)

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Something when wrong while querying")
	}

	return aboutMe, nil
}

func (s *aboutMeSettingService) DtoToRaw(settingDto *model.NewAboutMe) *model.AboutMe {
	aboutMe := model.AboutMe{ID: settingDto.ID, Name: settingDto.Name, Description: settingDto.Description, Content: settingDto.Content, ImgUrl: settingDto.ImgURL, SettingID: settingDto.SettingID}
	return &aboutMe
}
